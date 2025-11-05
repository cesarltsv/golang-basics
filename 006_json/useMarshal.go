package main

import (
	"encoding/json"
	"fmt"
)

func useMarshel() {
	response, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshel method: %v \n ", string(response))
}