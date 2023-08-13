package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.ListenAndServe(":4000", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// gettin the query params of form values when it is been passed through the url, when no value is passed it returns an empty string
	v := req.FormValue("first_name")
	fmt.Fprint(w, "My first name is:"+v)
}
