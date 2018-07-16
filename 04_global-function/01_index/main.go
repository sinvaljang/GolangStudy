/*
   Global Functionであるindexを使う方法は以下のようだ
   index . [index]

   ex)
   main.go
       temp := []int {3,2,1,0}
   template
       {{index . 0}}
       -> 3
       {{index . 1}}
       -> 2
*/

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
	nums := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Junban []string
		Name   string
	}{
		nums,
		"fuzimoto",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
