package calculation

import "fmt"

func SumNumbers(numbers ...int) int {
	var result int
	for _, value := range numbers {
		result += value
	}
	printResult(result)
	return result
}

func printResult(result int) {
	fmt.Printf("The result is: %d \n", result)
}