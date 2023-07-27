package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p *person) multiplyAge() int {
	return p.Age * 3
}
func (p *person) doubleAge(n int) int {
	return p.Age * 2
}

func (p *person) ageNumber() int {
	return p.Age
}

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("methodSample.gohtml"))
}

func main() {
	p := person{
		"Kay Araba",
		27,
	}
	err := temp.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal(err)
	}
}
