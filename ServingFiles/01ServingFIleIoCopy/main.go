package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", serveOnlineImage)
	http.HandleFunc("/localImage", serveLocalImage)
	http.ListenAndServe(":4000", nil)
}

func serveOnlineImage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="https://avatars.githubusercontent.com/u/78743753?v=4">`)
}

func serveLocalImage(w http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "text/html")
	localFile, err := os.Open("SampleImage.jpeg")
	if err != nil {
		http.Error(w, "File not found", 404)
	}
	defer localFile.Close()
	// io.WriteString(w, `<img src="https://avatars.githubusercontent.com/u/78743753?v=4">`)
	io.Copy(w, localFile)
}
