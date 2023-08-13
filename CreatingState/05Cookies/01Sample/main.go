package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", WriteCookie)
	http.HandleFunc("/read-cookie", ReadCookies)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func WriteCookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "kay-sample-cookie",
		Value: "kay araba running go",
	})
	fmt.Fprintln(w, "Cokkie written successfully check browser devtools")
	fmt.Fprintln(w, "Check devtools application/cookies to locate cookies")
}

func ReadCookies(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("kay-sample-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "This is the cookie you requested for", c)
}
