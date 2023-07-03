package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/pos/models"
	"github.com/pos/utils"
)

func Food(w http.ResponseWriter, req *http.Request) {
	var foods models.Food
	if req.Method == http.MethodPost {

		data, _ := io.ReadAll(req.Body)
		json.Unmarshal(data, &foods)
		err := utils.Validate(foods)

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
			db.Model(&foods).Where("food_name=?", foods.Food_Name).Count(&count)
			if count >= 1 {
				respones := map[string]interface{}{
					"message": "already item availabel",
					"status":  http.StatusBadRequest,
				}
				response_json, _ := json.Marshal(respones)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response_json)
			} else {
				db.Create(&foods)
				respones := map[string]interface{}{
					"message": "sussces create item",
					"status":  http.StatusAccepted,
				}
				response_json, _ := json.Marshal(respones)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusAccepted)
				w.Write(response_json)
			}
		}

	} else if req.Method == http.MethodGet {

		var data []models.Food
		db := models.Getdatabase().Db
		db.Find(&foods).Scan(&data)
		respones := map[string]interface{}{
			"data":   data,
			"status": http.StatusOK,
		}
		response_json, _ := json.Marshal(respones)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response_json)
	}
}

func Variant(w http.ResponseWriter, req *http.Request) {
	var data []utils.Variant
	if req.Method == http.MethodGet {
		db := models.Getdatabase().Db
		db.Find(&data)
		respones := map[string]interface{}{
			"data":   data,
			"status": http.StatusOK,
		}
		response_json, _ := json.Marshal(respones)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response_json)
	}
}

func DeleteVariant(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodDelete {
		id, _ := strconv.Atoi(req.URL.Query().Get("id"))

		db := models.Getdatabase().Db
		db.Delete(&models.Variant{}, id)
		respones := map[string]interface{}{
			"message": "delete success",
			"status":  http.StatusOK,
		}
		response_json, _ := json.Marshal(respones)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response_json)
	}
}
