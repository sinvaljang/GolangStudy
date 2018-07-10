/*
sliceをtemplateに
*/
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

/* 初期化関数 */
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	/* Slice */
	gods := []string{"buddha", "Jesus", "Muhammad"}

	/* Sliceをtemplateに渡す*/
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", gods)
	if err != nil {
		log.Fatalln(err)
	}
}
