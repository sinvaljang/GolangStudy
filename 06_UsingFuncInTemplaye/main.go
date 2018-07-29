package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) PrintSome() int {
	return 10
}

func (p person) AgeDouble() int {
	return (p.Age * 2)
}

func (p person) AgeByArgs(x int) int {
	return (p.Age * x)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./01_template/tpl.gohtml"))
}

func main() {
	p := person{
		Name: "fuzimoto",
		Age:  24,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
