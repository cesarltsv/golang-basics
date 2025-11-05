package main

import (
	"fmt"
	"io"
	"net/http"
)

func basicRequest() {
	request, err := http.Get("https://viacep.com.br/ws/01001000/json/")
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()
	defer fmt.Println("final line")
	fmt.Println("sasukee")
	res, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	} 
	fmt.Printf("%v \n", string(res))
	
}