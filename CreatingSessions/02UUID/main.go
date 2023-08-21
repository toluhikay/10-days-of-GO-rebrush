package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/get-uuid", getUUIDSession)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/favicon.ico", http.NotFound)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func getUUIDSession(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		uuid := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session",
			Value:    uuid.String(),
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Fprint(w, cookie.Name, ":", cookie.Value)
}
func homePage(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/get-uuid", http.StatusSeeOther)
	}

	fmt.Fprint(w, "Welcome to the home page for generating UUID", cookie)
}
