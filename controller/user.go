package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pos/models"
	"github.com/pos/utils"
)

func Migrate(w http.ResponseWriter, req *http.Request) {
	connection_db := models.Getdatabase()
	db := connection_db.Db

	db.AutoMigrate(models.User{}, models.Sift{}, models.Variant{}, models.Food{}, models.Sift{}, models.Transaction{}, models.Item_transaction{})
	// db.AutoMigrate(models.Item_transaction{})

	fmt.Println("ok")
}
func Register(w http.ResponseWriter, req *http.Request) {
	var register_data models.User
	// var error_handle []utils.HandleErrror
	if req.Method == http.MethodPost {
		data, _ := io.ReadAll(req.Body)

		json.Unmarshal(data, &register_data)
		pass := utils.PassswordHash(register_data.Password)
		register_data.Password = string(pass)
		err := utils.Validate(register_data)

		if err != nil {
			error_message := map[string]interface{}{
				"error":  err,
				"status": http.StatusBadRequest,
			}

			data, _ := json.Marshal(error_message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(data)
		} else {
			var count int64

			db := models.Getdatabase().Db

			db.Model(&register_data).Where("email=?", register_data.Email).Count(&count)
			if count >= 1 {
				fmt.Println("already availabel account")
				respones := map[string]interface{}{
					"message": "user has availabel",
					"status":  http.StatusBadRequest,
				}
				response_json, _ := json.Marshal(respones)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response_json)
			} else {
				db.Create(&register_data)
				respones := map[string]interface{}{
					"message": "sussces create user",
					"status":  http.StatusOK,
				}
				response_json, _ := json.Marshal(respones)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusAccepted)
				w.Write(response_json)
			}

		}

	}
}
func Login(w http.ResponseWriter, req *http.Request) {
	var login utils.Login
	var user models.User
	if req.Method == http.MethodPost {
		data, _ := io.ReadAll(req.Body)
		json.Unmarshal(data, &login)
		err := utils.Validate(login)

		if err != nil {
			error_message := map[string]interface{}{
				"error":  err,
				"status": http.StatusBadRequest,
			}

			data, _ := json.Marshal(error_message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(data)
		} else {
			db := models.Getdatabase().Db
			check := db.Where("email = ? ", login.Email).First(&user)
			if check.RowsAffected > 0 {
				compare_pass := utils.ComparePassword(user.Password, login.Password)
				if compare_pass {
					token, _ := utils.CreateCredentials(user.Email, user.Password)
					response := map[string]interface{}{
						"greeting": "Hello " + user.Fullname + "  you get the token, please check in bellow",
						"token":    token,
						"status":   http.StatusOK,
					}

					response_json, _ := json.Marshal(response)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(response_json)
				} else {
					error_message := map[string]interface{}{
						"error":  "password wrong",
						"status": http.StatusBadRequest,
					}

					data, _ := json.Marshal(error_message)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					w.Write(data)
				}
			} else {
				error_message := map[string]interface{}{
					"error":  "email worng",
					"status": http.StatusBadRequest,
				}

				data, _ := json.Marshal(error_message)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(data)
			}
		}
	}
}
