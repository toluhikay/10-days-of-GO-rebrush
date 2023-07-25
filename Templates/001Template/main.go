package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl template.Template

func init() {
	tpl = *template.Must(template.ParseFiles("templates/html.gohtml"))
}

func main() {
	// tpl, err := template.ParseGlob("templates/*")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Kay Araba")

}
