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
	Age  int
}

type teacher struct {
	Name    string
	Subject string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"mf": customFunc,
}

func init() {
	//FUNCMAP ADD
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
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
		Name: "nakamura",
		Age:  15,
	}

	fuzimoto := student{
		Name: "fuzimoto",
		Age:  15,
	}

	wada := teacher{
		Name:    "wada",
		Subject: "computer",
	}

	hitomi := teacher{
		Name:    "hitomo",
		Subject: "c",
	}

	students := []student{nakamura, fuzimoto}

	teachers := []teacher{wada, hitomi}

	/*structの変数名は大文字にすることを注意すること*/
	data := struct {
		St []student
		Te []teacher
	}{
		students,
		teachers,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
