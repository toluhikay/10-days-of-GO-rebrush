package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("predefinedGlb.gohtml"))
	newFile, err := os.Create("predefinedGlb.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	g1 := struct {
		Score1 int
		Score2 int
	}{
		20,
		200,
	}

	temp.ExecuteTemplate(newFile, "predefinedGlb.gohtml", g1)
}

func main() {

}
