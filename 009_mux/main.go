package main

import "net/http"

type home struct {
	title string
}

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.title))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/home", home{ title: "First Mux"})
	http.ListenAndServe(":8080", mux)

	// mux2 := http.NewServeMux()
	// mux2.Handle("/home", home{ title: "Secconde Mux"})
	// http.ListenAndServe(":8081", mux2)
}

