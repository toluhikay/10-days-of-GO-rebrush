package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.HandlerFunc(HandReq))
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func HandReq(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Fprint(w, "This is KAy testing the http notfound handler")
}
