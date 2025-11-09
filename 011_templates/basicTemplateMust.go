package main

import (
	"fmt"
	"html/template"
	"os"
)

func useBasicTemplateMust() {
	course := Course{
		Name: "Must template",
		Hours: 123,
	}
	t := template.Must(template.New("Course template").Parse("Course: {{ .Name }} - Hours: {{ .Hours }}\n"))
	err := t.Execute(os.Stdout, course)
	if err != nil {
		fmt.Println(err)
	}
}