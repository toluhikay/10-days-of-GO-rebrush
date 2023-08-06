package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", serveContent)
	http.HandleFunc("/serveFile", serveFile)
	http.ListenAndServe(":4000", nil)
}

// serving file with ServeContent
func serveContent(w http.ResponseWriter, req *http.Request) {
	// step one open the file
	file, err := os.Open("SampleImage.jpeg")
	if err != nil {
		http.Error(w, "File not FOund", 404)
	}
	defer file.Close()

	// get the details of the file using Stat method available to File
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "File not Found", 404)
	}

	http.ServeContent(w, req, file.Name(), fileInfo.ModTime(), file)

}

// serving file with serve file
func serveFile(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("SampleImage.jpeg")
	if err != nil {
		http.Error(w, "File not Found", 404)
	}
	defer file.Close()

	http.ServeFile(w, req, file.Name())
}
