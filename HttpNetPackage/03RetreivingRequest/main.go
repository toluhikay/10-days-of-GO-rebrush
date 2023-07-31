package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

// trying to get the request bodies
type httpResponseWriter int

func (wr httpResponseWriter) ServeHTTP(w http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	if err != nil{
		log.Fatal(err)
	}
	data := struct{
		Method string
		Url *url.URL
		Submissions map[string][]string
		Header http.Header
		Host string
		ContentLength int
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		int(req.ContentLength),
	}
	temp.ExecuteTemplate(w, "request.gohtml", data)
}

var temp template.Template

func init(){
	temp = *template.Must(template.ParseFiles("request.gohtml"))
}



func main(){
	var d httpResponseWriter
	http.ListenAndServe(":4000", d)
}