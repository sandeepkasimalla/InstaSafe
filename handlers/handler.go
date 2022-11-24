package handlers

import (
	"InstaSafe/common"
	validator "InstaSafe/payloadvalidator"
	"InstaSafe/service"
	"encoding/json"
	"net/http"

	"github.com/go-chassis/openlog"
)

type Handler struct {
	Service service.Service
}

type Response struct {
	Msg    string      `json:"_msg"`
	Status int         `json:"_status"`
	Data   interface{} `json:"data"`
}

func (h *Handler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	openlog.Info("Got a request to add Transaction")
	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&data)
	valres, err := validator.ValidatePaylaod("./../payloadschemas/transaction.json", data)
	if err != nil {
		openlog.Error(err.Error())
		response := Response{Msg: err.Error(), Data: valres, Status: 400}
		json.NewEncoder(w).Encode(response)
		return
	}
	input := common.AddTransactionInput{Data: data}
	res := h.Service.AddTransaction(input)
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) FetchAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := h.Service.FetchAllTransactions()
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) DeleteAllTransactions(w http.ResponseWriter, r *http.Request) {
	openlog.Info("Got a request to Delete All Transactions")
	// set header.
	w.Header().Set("Content-Type", "application/json")
	res := h.Service.DeleteAllTransactions()
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res)
}
