package main

import (
	"io"
	"log"
	"net/http"
)

type httpRequestWriter int

// handling of routes with Vanilla Go
func (h httpRequestWriter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "This is the dog route")
	case "/cat":
		io.WriteString(w, "this is the cat route")
	}
}

func main() {
	var d httpRequestWriter
	err := http.ListenAndServe(":4000", d)
	if err != nil {
		log.Fatal(err)
	}

}
