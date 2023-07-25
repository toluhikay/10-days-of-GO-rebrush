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
	newFile, err := os.Create("variableSample.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	templ.Execute(newFile, "Ich bin Kay und ich bin sehr klug - sample of the german i learnt")
}
