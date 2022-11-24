package service

import (
	"InstaSafe/common"
	"fmt"
	"time"
)

type Service struct {
}

func (h *Service) AddTransaction(input common.AddTransactionInput) common.Response {
	timestamp := input.Data["timestamp"].(string)
	transactiontime, err := time.Parse("2006-01-02T15:04:05.999Z", timestamp)
	if err != nil {
		fmt.Println(err)
		return common.Response{Msg: "Internal Server Error", Data: nil, Status: 500}
	}
	curtime := time.Now()
	diff := curtime.Sub(transactiontime)

	if diff > time.Duration(time.Second*60) {
		fmt.Println("duration :", diff)
		return common.Response{Msg: "Timestamp shoud be with 60 seconds", Data: input.Data, Status: 204}
	} else if diff < 0 {
		return common.Response{Msg: "Timestamp shoud not be in future", Data: input.Data, Status: 422}
	}

	common.Transactions = append(common.Transactions, input.Data)
	return common.Response{Msg: "Transaction added successfully", Data: input.Data, Status: 201}
}

func (h *Service) FetchAllTransactions() common.Response {

	stats, err := common.GetStats(time.Now())
	if err != nil {
		fmt.Println(err)
		return common.Response{Msg: "Internal Server Error", Data: nil, Status: 500}
	}
	return common.Response{Msg: "All Transaction Fetched successfully", Data: stats, Status: 200}
}

func (h *Service) DeleteAllTransactions() common.Response {
	common.Transactions = nil
	return common.Response{Msg: "All Transactions deleted successfully", Data: nil, Status: 204}
}
