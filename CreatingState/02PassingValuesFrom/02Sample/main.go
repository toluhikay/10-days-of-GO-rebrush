package main

import (
	"html/template"
	"log"
	"net/http"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	f := req.FormValue("first_name")
	l := req.FormValue("last_name")
	s := req.FormValue("subscribed") == "on"

	err := temp.ExecuteTemplate(w, "index.gohtml", person{f, l, s})
	if err != nil {
		log.Fatal(err)
	}

}
