package main

import (
	"fmt"
	"net/http"
)

// we will see the differenc between a Handle and a handleFUnc methods of the http pakaage

func dogRoute(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Knowing the difference between a Handle method and HanldeFUnc method")
}
func dogRoute2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Knowing the difference between a Handle method and HanldeFUnc method dog route2")
}
func catRoute(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Knowing the difference between a Handle method and HanldeFUnc method 1 cat route")
}
func catRoute2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Knowing the difference between a Handle method and HanldeFUnc method 2 cat route")
}

func main() {

	// using the handle method which takes in a pattern and a http Handler method
	http.Handle("/dog1", http.HandlerFunc(dogRoute))
	http.Handle("/cat1", http.HandlerFunc(catRoute))

	// using the HandleFUnc Method that takes in a pattern and a function that receieves responsewriter and *request as it parameters
	http.HandleFunc("/dog2", dogRoute2)
	http.HandleFunc("/cat2", catRoute2)

	http.ListenAndServe(":4000", nil)

}
