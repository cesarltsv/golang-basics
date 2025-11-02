package main

import "fmt"

func useTypeAssertion() {
	var value interface{} = "teste"
	fmt.Println(value.(string))
	res, ok := value.(int)
	fmt.Printf("value: %v * result: %v \n", res, ok)

}