package main

import "fmt"

func useSlices() {
	myFirstSlice := []int{ 12, 24, 34, 46, 213, 8678, 678, 456, 6456}
	fmt.Printf("len=%d cap=%d %v \n", len(myFirstSlice), cap(myFirstSlice), myFirstSlice)
	fmt.Printf("len=%d cap=%d %v \n", len(myFirstSlice[:0]), cap(myFirstSlice[:0]), myFirstSlice[:0])
	fmt.Printf("len=%d cap=%d %v \n", len(myFirstSlice[:4]), cap(myFirstSlice[:4]), myFirstSlice[:4])
	fmt.Printf("len=%d cap=%d %v \n", len(myFirstSlice[2:]), cap(myFirstSlice[2:]), myFirstSlice[2:])

	myFirstSlice = append(myFirstSlice, 999)
	fmt.Printf("len=%d cap=%d %v \n", len(myFirstSlice), cap(myFirstSlice), myFirstSlice)

}