package main

import "fmt"

type CustomId float64

var userId CustomId
var productName string

func createTypes() {
	userId = 123.23
	fmt.Printf("CustomId: %f \n", userId)
	fmt.Printf("productName type: %T \n", productName)
	fmt.Printf("userId type: %T \n", userId)
}