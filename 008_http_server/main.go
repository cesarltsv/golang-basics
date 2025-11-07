package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/", SearchZipcode)
	http.ListenAndServe(":8080", nil)
}

func SearchZipcode(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("path %s \n", r.URL.Path)
	cepParam := r.URL.Query().Get("zipcode")
	if len(cepParam) < 7 || cepParam == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("address invalid"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello"))
}