package main

import (
	"fmt"
	"net/http"
)

// here I will be using the Hnadler func which takes in a pattern and a hndler func that takes in parameter of http>ResponseWrite and *http.Request

func dogRoute(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Mein hund ist super: meaning my dog is great in english</h1>")
}

func catRoute(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "<p>Meine katze is schon: meaning in english my cat is pretty</p>")
}

func main() {
	http.HandleFunc("/dog/", dogRoute)
	http.HandleFunc("/cat", catRoute)

	http.ListenAndServe(":4000", nil)
}
