package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type student struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	/* 学生情報１ */
	nakamura := student{
		Name: "nakamura",
		Age:  14,
	}

	/* 学生情報2 */
	fuzimoto := student{
		Name: "fuzimoto",
		Age:  15,
	}

	/* 学生情報を保存するslice */
	students := []student{nakamura, fuzimoto}

	err := tpl.Execute(os.Stdout, students)
	if err != nil {
		log.Fatalln(err)
	}
}
