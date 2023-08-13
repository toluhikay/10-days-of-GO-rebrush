package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// writeing an uploaded file to a new file
func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", CreateNewFile)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func CreateNewFile(w http.ResponseWriter, req *http.Request) {
	// step 1 ; create a string to read the file contents to
	var s string
	// step 2 check if the method of submitting the file is method post
	if req.Method == http.MethodPost {
		// step 3 open the file
		f, h, err := req.FormFile("file_upload")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer f.Close()
		fmt.Println("\n file:", f, "\n header:", h)

		// step 4 read the file content
		bs, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//step 5 convert the bytes gotten from reading the file into strings
		s = string(bs)

		// step 6 create a new file to read the uploaded file to
		newFile, err := os.Create(filepath.Join("./files/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer newFile.Close()

		// step 7 write the contents gotten to the new file
		_, err = newFile.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	// write back to the request
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<form action="" method="post" enctype="multipart/form-data">
	<input type="file" name="file_upload"/> <br />
	<input type="submit" name="upload" />
	</form>`+s)
}
