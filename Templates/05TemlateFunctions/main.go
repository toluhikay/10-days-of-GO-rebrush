package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fn = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fn).ParseFiles("sampleFunc.gohtml"))
}

type sage struct {
	Name string
}

func main() {
	newFile, err := os.Create("SampleFunc.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	a := sage{
		Name: "Kay Araba",
	}
	b := sage{
		Name: "Kay Araba",
	}
	c := sage{
		Name: "Kay Araba",
	}

	sages := []sage{a, b, c}
	fmt.Println(sages)

	tpl.ExecuteTemplate(newFile, "sampleFunc.gohtml", sages)
}
