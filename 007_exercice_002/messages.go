package main

import "fmt"


func showResult(address Address) {
	fmt.Printf("\n ------> Your Address: \n")
	fmt.Printf("\t State: %s - (%s) \n", address.State, address.StateId)
	fmt.Printf("\t Street: %s \n", address.Street)
	fmt.Printf("\t Neighbourhood: %s \n", address.Neighbourhood)
	fmt.Printf("\t City: %s \n", address.Location)
}

func errorToGetResult() {
	fmt.Printf("\n******* ERROR *******\n")
	fmt.Printf("zipcode: %s not found! ", Zipcode)
}