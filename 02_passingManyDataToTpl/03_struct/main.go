/*
How to pass Struct into Template
*/

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

/* 学生の情報をStructに定義する */
type student struct {
	Name  string
	Class string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	/* Struct init */
	nakamura := student{
		Name:  "nakamura",
		Class: "1class",
	}

	/* passing struct into template*/
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nakamura)
	if err != nil {
		log.Println(err)
	}
}
