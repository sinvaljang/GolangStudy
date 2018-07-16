package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Age   int
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	user1 := user{
		Name:  "Fuzimoto",
		Age:   24,
		Admin: true,
	}

	user2 := user{
		Name:  "Jang",
		Age:   30,
		Admin: false,
	}

	user3 := user{
		Name:  "",
		Age:   6,
		Admin: true,
	}

	users := []user{user1, user2, user3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
