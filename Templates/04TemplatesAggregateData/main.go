package main

import (
	"log"
	"os"
	"text/template"
)

var temp template.Template

type office struct {
	Office string
	Number int
}

func init() {
	temp = *template.Must(template.ParseGlob("templates/*"))
}

func main() {

	newFile, err := os.Create("structSampleResult.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// slice sample
	// sages := []string{"Kay", "Araba", "Damoche"}

	// struct sample
	arabaOffice := office{
		Office: "Software Developer",
		Number: 2,
	}

	temp.ExecuteTemplate(newFile, "structSample.gohtml", arabaOffice)

}
