package main

import (
	"fmt"
	"net/http"
)

type test string

func (t test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	var d test
	http.ListenAndServe(":8080", d)
}
