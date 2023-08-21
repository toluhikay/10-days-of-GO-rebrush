package main

import "net/http"

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	// check the session id to return the value which is the user name stored in it
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
