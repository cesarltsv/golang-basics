package main

import "net/http"


func main() {
	http.HandleFunc("/", SearchZipcode)
	http.ListenAndServe(":8080", nil)
}

func SearchZipcode(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}