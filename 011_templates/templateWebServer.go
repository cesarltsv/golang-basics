package main

import (
	"html/template"
	"net/http"
)

func templateWebServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeController)
	http.ListenAndServe(":8080", mux)
}

func homeController(w http.ResponseWriter, r *http.Request) {
	template := getTemplate();

	err := template.Execute(w, Courses{
		{"Go", 40},
		{"Javascript", 40},
		{"C#", 40},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}

func getTemplate() *template.Template {
	templates := []string{
		"template.html",
		"table.html",
	}
	t := template.Must(template.New("template.html").ParseFiles(templates...))
	return t
}

