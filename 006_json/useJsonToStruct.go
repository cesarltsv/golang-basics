package main

import (
	"encoding/json"
	"fmt"
)

func useJsonToStruct() {
	jsonPure := []byte(`{"numero": 123123, "saldo": 999 }`)
	var account Account

	err := json.Unmarshal(jsonPure, &account)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Converted json to struct: %v \n", account)
	fmt.Printf("Number: %v \n", account.Number)
	fmt.Printf("Amount: %v \n", account.Amount)
}