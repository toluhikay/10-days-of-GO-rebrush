package main

import (
	"fmt"
	"log"
	"net/http"
)

type httpSample int

func (d httpSample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Implementing the Handler interface with my own type")
}

func main() {
	var connectRequest httpSample
	err := http.ListenAndServe(":4000", connectRequest)
	if err != nil {
		log.Fatal(err)
	}
}
