package main

import "fmt"

func useArrays() {
	var myFirstArray [3]int
	myFirstArray[0] = 20
	myFirstArray[1] = 10
	fmt.Printf("Array size: %d \n", len(myFirstArray))
	fmt.Printf("Array position 0: %d \n", myFirstArray[0])
	fmt.Printf("Array position 1: %d \n", myFirstArray[1])
	fmt.Printf("Array position 2: %d \n", myFirstArray[2])

	fmt.Println("\n **** Loop a array ****")
	for i, v := range myFirstArray {
		fmt.Printf("Value of position %d, is:  %d \n", i, v)
	}
}