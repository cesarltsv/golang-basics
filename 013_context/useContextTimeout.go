package main

import (
	"log"
	"net/http"
	"time"
)


func useContextTimeout() {
	http.HandleFunc("/", handlerController)
	http.ListenAndServe(":8080", nil)
}

func handlerController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciado")
	defer log.Println("Request finalizado")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request process!")
		w.Write([]byte("Request process!"))
	case <-ctx.Done():
		log.Println("request canceled by client")
		http.Error(w, "request canceled by client", http.StatusRequestTimeout)
	}
}