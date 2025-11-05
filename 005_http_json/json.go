package main

import "encoding/json"
type Account struct {
	Id string
	Name string
	Amount int
}

func transformStructToJason() {
	account := Account{
		Id: "1233",
		Name: "Itadori Iuji",
		Amount: 342,
	}

	data, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}

	println(string(data))
}