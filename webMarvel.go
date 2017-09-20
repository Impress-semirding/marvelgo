package main

import (
	"net/http"
	"io"
	"log"
)

func helloHandler(w http.ResponseWriter, request *http.Request)  {
	io.WriteString(w, "<html><body><form method=\"POST\" action=\"/upload\" "+
		" enctype=\"multipart/form-data\">"+
		"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
		"<input type=\"submit\" value=\"Upload\" />"+
		"</form></body></html>")
}

func main() {

	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("有错误")
	}
}

