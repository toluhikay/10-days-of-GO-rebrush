package main

import (
	"log"
	"os"
	"text/template"
)

var templ *template.Template

var fn = template.FuncMap{
	"squareNum": squareNum,
	"doubleNum": doubleNum,
}

func init() {
	templ = template.Must(template.New("pipeline.html").Funcs(fn).ParseFiles("pipeline.gohtml"))
}

func squareNum(n float64) float64 {
	n = n * n
	return n
}

func doubleNum(n float64) float64 {
	return n * 2
}

func main() {
	newFile, err := os.Create("pipeline.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	templ.ExecuteTemplate(newFile, "pipeline.gohtml", float64(5))
}
