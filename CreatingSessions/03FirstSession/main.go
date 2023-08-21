package main

import (
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName  string
	FirstName string
	LastName  string
}

var temp template.Template

// create a map that stores the user id and the user
var dbUsers = make(map[string]user)

// create a map that stores session id with user id as the value
var dbSessions = map[string]string{}

func init() {
	temp = *template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", createSession)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/get-user", getUser)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func createSession(w http.ResponseWriter, req *http.Request) {
	// first check if there is a session id in the cookie, if not then create one
	cookie, err := req.Cookie("session")
	if err != nil {
		sId := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(w, cookie)
	}

	// if user already exists then get user
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	// if no user then submit the user form
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		u = user{un, fn, ln}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}

	temp.ExecuteTemplate(w, "index.gohtml", u)

}

func getUser(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// get the user
	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]

	temp.ExecuteTemplate(w, "bar.gohtml", u)

}
