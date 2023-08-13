package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; chatsey=utf-8")
	v := req.FormValue("first_name")
	io.WriteString(w, `<form method="GET">
	<input type="text" name="first_name" />
	<input type="submit" />
	</form>`+v)
}
