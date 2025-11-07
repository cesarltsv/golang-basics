package main

import (
	"io"
	"net/http"
)


func searchAddress(zipcode string) []byte {
url := "http://viacep.com.br/ws/" + zipcode + "/json/"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	res, err  := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return res
}