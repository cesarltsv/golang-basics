package main

type Course struct {
	Name string
	Hours int
}

type Courses []Course

func main(){
	useBasicTemplate()
	useBasicTemplateMust()
	useExternalTemplates()
}
