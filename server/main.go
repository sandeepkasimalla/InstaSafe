package main

import (
	"log"
	"net/http"

	"InstaSafe/handlers"

	"github.com/go-chassis/openlog"
	"github.com/gorilla/mux"

	"InstaSafe/service"
)

func GetService(dbname string) service.Service {
	return service.Service{}
}
func main() {

	r := mux.NewRouter()

	service := GetService("Users")
	h := handlers.Handler{Service: service}

	r.HandleFunc("/transactions", h.AddTransaction).Methods("POST")
	r.HandleFunc("/transactions", h.DeleteAllTransactions).Methods("DELETE")
	r.HandleFunc("/statistics", h.FetchAllTransactions).Methods("GET")
	openlog.Info("Started listening at http://localhost:8070")
	log.Fatal(http.ListenAndServe(":8070", r))

}
