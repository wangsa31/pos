package middelware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pos/utils"
)

func AuthMiddelware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			error_message := map[string]interface{}{
				"error":  "Token is required",
				"status": http.StatusUnauthorized,
			}
			responese_json, _ := json.Marshal(error_message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(responese_json)
			return
		}

		token_string := strings.Split(authHeader, "Bearer ")
		if len(token_string) != 2 {
			error_message := map[string]interface{}{
				"error":  "Invalid token format",
				"status": http.StatusUnauthorized,
			}
			responese_json, _ := json.Marshal(error_message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(responese_json)
			return
		}
		fmt.Println(token_string[1])

		valid, err := utils.CheckCredentials(token_string[1])

		if err != nil || !valid {
			error_message := map[string]interface{}{
				"error":  "Invalid token",
				"status": http.StatusUnauthorized,
			}
			responese_json, _ := json.Marshal(error_message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(responese_json)
			fmt.Println(err)
			return
		}

		next(w, req)
	}
}
