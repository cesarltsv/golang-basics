package main

import "fmt"

type GenericNumber interface {
	~int | ~float64
}


func useGenerics() {
	values := map[string]int{"bread": 123, "onion": 20, "potato": 24}
	result := handleCalculation(values)
	fmt.Printf("result is: %d \n", result)

	scores := map[string]float64{"bread": 5.45, "onion": 5.64, "potato": 2.35}
	resultScores := handleCalculation(scores)
	fmt.Printf("result is: %v \n", resultScores)
	fmt.Printf("Compared is 10 vs 10: %v \n", compare(10,10))
	fmt.Printf("Compared is 10 vs 15: %v \n", compare(10,15))


}

func handleCalculation[T GenericNumber](m map[string]T) T {
	var result T
	for _, value := range m {
		result += value
	}
	return result
}

func compare[T comparable](a, b T) bool {
	return a == b 
}