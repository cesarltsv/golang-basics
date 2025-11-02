package main

import "fmt"

const MAX_LENTH = 100
var machineName = "Jarvis"
var ( 
	isActive bool = true
	age int 
	score float64
)


func basicVariables() {
	var lastName string
	name := "Naruto"
	lastName = "Uzumaki"

	fmt.Printf("firstName %s \n", name)
	fmt.Printf("LastName %s \n", lastName)
	fmt.Printf("machine name %s \n", machineName)
	fmt.Printf("isActive: %t, age: %d, score: %f \n", isActive, age, score)

	fmt.Printf("calculation: %d \n", AddNumber(3, 4))
}