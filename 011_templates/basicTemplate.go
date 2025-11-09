package main

import (
	"fmt"
	"html/template"
	"os"
)

func useBasicTemplate() {
	course := Course{
		Name: "Basic template",
		Hours: 123,
	}

	temp := template.New("Course template")
	temp, _ = temp.Parse("Course: {{ .Name }} - Hours: {{ .Hours }} \n")
	err := temp.Execute(os.Stdout, course)
	if err != nil {
		fmt.Println(err)
	}
}