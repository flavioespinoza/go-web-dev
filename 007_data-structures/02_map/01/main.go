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

	eats := map[string]string{
		"American": "Cheese Burger and Fries",
		"Indian":   "Tiki Masala",
		"Japanese": "Yaki Soba",
		"Mexican":  "Chile Verde",
		"Russian":  "Vodka",
	}

	err := tpl.Execute(os.Stdout, eats)
	if err != nil {
		log.Fatalln(err)
	}
}
