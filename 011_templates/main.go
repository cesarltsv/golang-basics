package main

import (
	"fmt"
	"html/template"
	"os"
)

type Course struct {
	Name string
	Hours int
}

func main(){
	course := Course{
		Name: "name text",
		Hours: 123,
	}

	temp := template.New("Course template")
	temp, _ = temp.Parse("Course: {{ .Name }} - Hours: {{ .Hours }}")
	err := temp.Execute(os.Stdout, course)
	if err != nil {
		fmt.Println(err)
	}
}
