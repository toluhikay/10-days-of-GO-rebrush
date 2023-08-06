package main

import (
	"io"
	"net/http"
)

func main() {
	// implementing the http.fileserver
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/kay", kay)
	http.ListenAndServe(":4000", nil)

}

func kay(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="SampleImage.jpeg">`)
}
