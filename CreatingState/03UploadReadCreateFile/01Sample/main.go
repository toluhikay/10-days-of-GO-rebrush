package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// sample example of uploading a file reading the content and displaying the content

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", ReadFile)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func ReadFile(w http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {
		// open the file
		file, header, err := req.FormFile("file_upload")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer file.Close()
		fmt.Println("\n file:", file, "\n header:", header, "\n error:", err)

		//read the file
		bs, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset= utf-8")
	_, err := io.WriteString(w, `<form action="" method="post" enctype="multipart/form-data">
	<input type="file" name="file_upload"/> <br />
	<input type="submit" name="upload" />
	</form>`+s)
	if err != nil {
		log.Fatal(err)
	}

}
