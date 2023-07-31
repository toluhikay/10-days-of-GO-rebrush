package main

import (
	"fmt"
	"net/http"
)

type httpWriterRequest int

func (wr httpWriterRequest) ServeHTTP(w http.ResponseWriter, req *http.Request){
	// first step to reading forms throught either the url or body which returns an error if any
	
	w.Header().Set("Kay-Code", "Rebrushed my skills with headers")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Kay Doing So Great with Gp</h1>")

	

}





func main(){
	var d httpWriterRequest
	http.ListenAndServe(":4000", d)
}	