package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Name  string
	Class int
}

type School struct {
	Students []Student
	Teachers []Teacher
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	nakamura := Student{
		Name: "nakamura",
		Age:  15,
	}

	fuzimoto := Student{
		Name: "fuzimoto",
		Age:  15,
	}

	wada := Teacher{
		Name:  "wada",
		Class: 1,
	}

	hitomi := Teacher{
		Name:  "hitomi",
		Class: 2,
	}

	Students_S := []Student{
		nakamura,
		fuzimoto}

	Teachers_S := []Teacher{
		wada,
		hitomi}

	higashiE := School{
		Students: Students_S,
		Teachers: Teachers_S,
	}

	err := tpl.Execute(os.Stdout, higashiE)
	if err != nil {
		log.Fatalln(err)
	}
}
