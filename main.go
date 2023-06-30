package main

import (
	"log"
	"net/http"

	"github.com/pos/controller"
	"github.com/pos/middelware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/register", controller.Register)
	mux.HandleFunc("/user/login", controller.Login)
	mux.HandleFunc("/foods", controller.Food)
	mux.HandleFunc("/foods/add", middelware.AuthMiddelware(controller.Food))
	mux.HandleFunc("/transactions/add", middelware.AuthMiddelware(controller.Transactions))
	mux.HandleFunc("/transactions/list/all", controller.Transactions)
	mux.HandleFunc("/transactions", controller.FindTransaction)
	mux.HandleFunc("/variant/list/all", controller.Variant)

	// mux.HandleFunc("/user/migrate", controller.Migrate)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Printf("Server Connecting..... , Please Access in url http://localhost:8080/")
	log.Fatal(server.ListenAndServe())
}
