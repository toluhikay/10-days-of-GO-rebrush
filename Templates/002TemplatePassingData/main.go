package main

import (
	"log"
	"os"
	"text/template"
)

var templ template.Template

func init() {
	templ = *template.Must(template.ParseGlob("templates/*"))
}

func main() {

	nf, err := os.Create("newData.html")
	if err != nil {
		log.Fatal(err)
	}

	defer nf.Close()
	templ.Execute(nf, "Araba Mother Nest")

}
