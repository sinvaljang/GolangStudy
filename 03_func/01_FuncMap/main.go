package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type student struct {
	Name string
	age  int
}

type teacher struct {
	Name string
	Kind string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"mf": customFunc,
}

func init() {
	//tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func customFunc(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}

	return s
}

func main() {
	nakamura := student{
		"nakamura",
		15,
	}

	fuzimoto := student{
		"fuzimoto",
		15,
	}

	wada := teacher{
		"wada",
		"computer",
	}

	hitomi := teacher{
		"hitomo",
		"c",
	}

	students := []student{nakamura, fuzimoto}

	teachers := []teacher{wada, hitomi}

	data := struct {
		st []student
		te []teacher
	}{
		students,
		teachers,
	}

	//err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
