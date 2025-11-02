package main

import "fmt"

func usePointers() {
	//Get address in memory
	name := "teste"
	println(&name)

	var cloneName *string = &name
	fmt.Printf("address: %v * value: %s * originalValue: %s \n", cloneName, *cloneName, name)

	*cloneName = "naruto"
	fmt.Printf("address: %v * value: %s * originalValue: %s \n", cloneName, *cloneName, name)

	name = "batatone"
	fmt.Printf("address: %v * value: %s * originalValue: %s \n", cloneName, *cloneName, name)

	fmt.Printf("\n **** Use pointer to change original value **** \n \n")
	numberOne := 10
	numberTwo := 30
	fmt.Printf("BEFORE: numberOne: %d * numberTwo: %d \n", numberOne, numberTwo)
	result := calculation(&numberOne, &numberTwo)
	fmt.Printf("resul: %d \n", result)
	fmt.Printf("AFTER: numberOne: %d * numberTwo: %d \n", numberOne, numberTwo)

}

func calculation(a, b *int) int {
	*a = 1000
	return *a + *b
}