package main

import (
	"errors"
	"fmt"
)

func useFunctions() {
	fmt.Printf("the sum of 2 + 5 is:  %d \n", sumNumbers(2, 5))
	value, err := sumError(5, 2)
	if err != nil {
		fmt.Printf("error: %s \n", err)
	}
	fmt.Printf("the sum of function with error is:  %d \n", value)
	fmt.Printf("the variadict function result is :  %d \n", sumVariadic(4, 5, 5, 6, 7))

	closure := func() int {
		return sumVariadic(4, 5, 5, 6, 7) * 5
	}()

	fmt.Printf("the result of closure function is: %d \n", closure)
}

func sumNumbers(a int, b int) int {
	return a + b
}

func sumError(a, b int) (int, error) {
	if a + b <= 20 {
		return 0, errors.New("soma menor que 20")
	}
	return a + b , nil
}

func sumVariadic(numbers ...int) int {
	var result int
	for _, num := range numbers {
		result += num
	}
	return  result
}