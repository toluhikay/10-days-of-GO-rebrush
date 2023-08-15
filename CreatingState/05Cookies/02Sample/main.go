package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(SetCookie))
	http.HandleFunc("/delete", DeleteCookie)
	http.HandleFunc("/no-cookie", NoCookie)
	http.HandleFunc("/read-cookie", ReadCookie)
	http.HandleFunc("/favicon.ico", http.NotFoundHandler().ServeHTTP)
	http.ListenAndServe(":4000", nil)
}

func SetCookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Kay-cookie",
		Value: "kay-araba-encoded",
	})
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1> Kay Setting a new cookie</h1>`)
}

func ReadCookie(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Kay-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, c.String())
	fmt.Fprint(w, `<h1> Kay Setting a new cookie </h1>`, c.Value)
}

// deleting a cookie using maxAge

func DeleteCookie(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Kay-cookie")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, req, "/no-cookie", http.StatusSeeOther)
}

func NoCookie(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Check the cookie session you will see that there is no cookie")
}
