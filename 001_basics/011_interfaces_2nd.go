package main

import "fmt"


func useInterfaceSeconPart() {
	printValue(12)
	printValue("Teste")
	printValue(2.4444)

}

func printValue(value interface{}) {
	fmt.Printf("value: %v * type: %T \n", value, value)
}