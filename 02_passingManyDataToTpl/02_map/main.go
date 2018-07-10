package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	students := map[string]string{
		"1class": "nakamura",
		"2class": "kawamata",
		"3class": "fuzimoto"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", students)
	if err != nil {
		log.Fatalln(err)
	}
}
