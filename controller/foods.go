package controller

import (
	"encoding/json"
	"io"
	"net/http"

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
	var data []models.Variant
	if req.Method == http.MethodGet {
		db := models.Getdatabase().Db
		db.Find(&models.Variant{}).Scan(&data)
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
