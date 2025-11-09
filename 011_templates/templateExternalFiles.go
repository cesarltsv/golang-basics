package main

import (
	"html/template"
	"os"
)

func useExternalTemplates() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Javascript", 40},
		{"C#", 40},
	})

	if err != nil {
		panic(err)
	}
}