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

	horsemen := []string{"white", "red", "black", "pale"}

	err := tpl.Execute(os.Stdout, horsemen)
	if err != nil {
		log.Fatalln(err)
	}
}
