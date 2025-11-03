package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// creation	
	f, err := os.Create("users.txt")
	if err != nil {
		panic(err)
	}

	fileZise, err := f.WriteString("Heelo World")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created! size: %v \n", fileZise)
	f.Close()

	// read
	file, err := os.ReadFile("default.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", string(file))

	//Open in parts
	newFile, err := os.Open("largeText.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(newFile)
	bufferSize := make([]byte, 10)

	for {
		n, err := reader.Read(bufferSize)
		if err != nil{
			break
		}
		fmt.Println(string(bufferSize[:n]))
	}
}
