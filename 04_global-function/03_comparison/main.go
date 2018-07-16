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
	testRes := struct {
		Math int
		Sci  int
	}{
		80,
		90,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", testRes)
	if err != nil {
		log.Fatalln(err)
	}
}
