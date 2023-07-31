package main

import (
	"log"
	"net/http"
	"text/template"
)

type httpResponseWriter int

func (h httpResponseWriter) ServeHTTP(w http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	if err != nil{
		log.Fatal(err)
	}
	temp.ExecuteTemplate(w, "form.gohtml", req.Form)
}
var temp * template.Template

func init(){
	temp = template.Must(template.ParseFiles("form.gohtml"))
}
func main(){
	var d httpResponseWriter
	http.ListenAndServe(":4000", d)
}