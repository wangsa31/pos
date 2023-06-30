package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pos/models"
	"github.com/pos/utils"
)

func Transactions(w http.ResponseWriter, req *http.Request) {
	var transactions utils.Transaction
	var transactions_db models.Transaction
	var transactions_items_db models.Item_transaction
	if req.Method == http.MethodPost {
		data, _ := io.ReadAll(req.Body)
		json.Unmarshal(data, &transactions)

		err := utils.Validate(transactions)

		if err != nil {
			error_message := map[string]interface{}{
				"error":  err,
				"status": http.StatusBadRequest,
			}

			data, _ := json.Marshal(error_message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(data)
		}

		if transactions.Total > transactions.Amount {
			error_message := map[string]interface{}{
				"error":  "Apologies, but you have insufficient funds",
				"status": http.StatusBadRequest,
			}

			response_json, _ := json.Marshal(error_message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response_json)
		} else {
			db := models.Getdatabase().Db

			transactions_db.Id = transactions.Id_transaction
			transactions_db.Id_user = int(transactions.Id_user)
			transactions_db.Pay_methode = transactions.Method_pay
			transactions_db.Status = "success"
			transactions_db.Type = transactions.Type
			transactions_db.Total_amount = int(transactions.Total)
			db.Create(&transactions_db)
			for _, items := range transactions.Item {
				transactions_items_db.Id_transaction = transactions.Id_transaction
				transactions_items_db.Id_foods = int(items.Id_food)
				transactions_items_db.Id_variant = int(items.Id_variant)
				transactions_items_db.Qty = int(items.Qty)
				db.Create(&transactions_items_db)
			}
			response := map[string]interface{}{
				"message": "Payment successful. Your transaction has been processed successfull",
				"status":  http.StatusOK,
			}

			response_json, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(response_json)
		}

	} else if req.Method == http.MethodGet {
		var x []utils.Result
		db := models.Getdatabase().Db
		db.Model(&models.Transaction{}).Select("users.fullname, transactions.id, foods.price , transactions.pay_methode, transactions.status, transactions.type, foods.food_name, variants.variant, item_transactions.qty").Joins("INNER JOIN users ON users.id = transactions.id_user").Joins("INNER JOIN item_transactions ON item_transactions.id_transaction = transactions.id ").Joins("INNER JOIN foods ON item_transactions.id_foods = foods.id").Joins("INNER JOIN variants ON item_transactions.id_variant = variants.id").Scan(&x)
		respones := map[string]interface{}{
			"data":   x,
			"status": http.StatusOK,
		}
		response_json, _ := json.Marshal(respones)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response_json)
	}
}

func FindTransaction(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("faktur")
	if req.Method == http.MethodGet {
		var x []utils.Result
		db := models.Getdatabase().Db
		db.Model(&models.Transaction{}).Select("users.fullname, transactions.id, foods.price , transactions.pay_methode, transactions.status, transactions.type, foods.food_name, variants.variant, item_transactions.qty").Joins("INNER JOIN users ON users.id = transactions.id_user").Joins("INNER JOIN item_transactions ON item_transactions.id_transaction = transactions.id ").Joins("INNER JOIN foods ON item_transactions.id_foods = foods.id").Joins("INNER JOIN variants ON item_transactions.id_variant = variants.id").Where("transactions.id = ?", id).Scan(&x)
		respones := map[string]interface{}{
			"data":   x,
			"status": http.StatusOK,
		}
		response_json, _ := json.Marshal(respones)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response_json)
	}

}
