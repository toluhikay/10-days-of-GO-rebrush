package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/name", handleData)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleData(w http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("gotemplate.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	temp.ExecuteTemplate(w, "gotemplate.gohtml", "pro")
}
