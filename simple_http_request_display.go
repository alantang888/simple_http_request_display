package main

import (
	"net/http"
	"log"
	"fmt"
	"bytes"
)

func requestHandler (w http.ResponseWriter, r *http.Request){
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("Client from %v, request path: %v\n", r.RemoteAddr, r.URL.Path))
	for _, cookie := range r.Cookies(){
		buffer.WriteString(fmt.Sprintf("\tCOOKIE[\"%v\"]: %v\n", cookie.Name, cookie.Value))
	}
	message := buffer.String()
	log.Printf(message)
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
