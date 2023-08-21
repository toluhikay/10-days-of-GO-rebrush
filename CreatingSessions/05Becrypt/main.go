package main

import (
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName  string
	FirstName string
	LastName  string
	Password  []byte
}

var temp *template.Template

var dbUsers = map[string]user{}
var dbSessions = make(map[string]string)

func init() {
	temp = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", getUser)
	http.HandleFunc("/sign-up", createUser)
	http.HandleFunc("/favicon.ico", http.NotFound)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func createUser(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// if there is no user already logged in then you can now allow for user sign up
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		pw := req.FormValue("password")

		// check if the user name already exists
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "User already exists", http.StatusForbidden)
		}

		// if no user matches user name then create a session for new user
		sId := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(w, c)

		// store the user name to the session id
		dbSessions[c.Value] = un
		// bcrypt pass word

		bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		// store the user to database
		u := user{un, fn, ln, bs}
		dbUsers[un] = u

		// once sign up is successful redirect to the signed up page
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	temp.ExecuteTemplate(w, "signUp.gohtml", nil)
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
