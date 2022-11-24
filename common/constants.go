package common

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type Response struct {
	Msg    string      `json:"_msg"`
	Status int         `json:"_status"`
	Data   interface{} `json:"data"`
}

type AddTransactionInput struct {
	Data map[string]interface{}
}

var Transactions []map[string]interface{}

func GetStats(curtime time.Time) (map[string]interface{}, error) {
	var max, min, average, count, sum float64

	if len(Transactions) > 0 {
		min = math.MaxFloat64
	}
	for _, transaction := range Transactions {
		timestamp := transaction["timestamp"].(string)
		transactiontime, err := time.Parse("2006-01-02T15:04:05.999Z", timestamp)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		diff := curtime.Sub(transactiontime)
		if diff < time.Duration(time.Second*60) {
			amount := transaction["amount"].(string)
			cost, err := strconv.ParseFloat(amount, 64)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			sum += cost
			count++
			average = sum / count
			if cost > max {
				max = cost
			}
			if cost < min {
				min = cost
			}
		}
	}
	return map[string]interface{}{
		"sum":     sum,
		"average": average,
		"count":   count,
		"min":     min,
		"max":     max,
	}, nil
}
