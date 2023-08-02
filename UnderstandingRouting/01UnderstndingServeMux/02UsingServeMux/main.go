package main

import (
	"fmt"
	"net/http"
)

type dogRequest int
type catRequest int

func (d dogRequest) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset-utf-8")
	fmt.Fprintln(w, "<h1>This is Kay Revisiting serve mux type using the NewServeMux function</h1>")
}

func (c catRequest) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h2>This is the cat request route</h2>")
}

func main() {
	var d dogRequest
	var c catRequest

	// this is using the NewServeMux
	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	// comming down I can make the code look elegant so far there is defaultServeMux I can remove ths NewServeMux
	http.Handle("/dog", d)
	http.Handle("/cat/", c)

	// this if or using serveMux I can change to using default serve mux byt turning the handler to nil
	// http.ListenAndServe(":4000", mux)

	// using default serveMux
	http.ListenAndServe(":5000", nil)
}
