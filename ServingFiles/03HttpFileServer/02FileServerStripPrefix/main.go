package main

import (
	"io"
	"net/http"
)

func main() {
	// using the strip prefix to server files
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/kaypicture", kayPicture)
	http.ListenAndServe(":4000", nil)

}

func kayPicture(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/assets/SampleImage.jpeg">`)
}
