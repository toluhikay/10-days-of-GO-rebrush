package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := temp.ExecuteTemplate(os.Stdout, "sample.gohtml", 100)
	if err != nil {
		log.Fatal(err)
	}
}
