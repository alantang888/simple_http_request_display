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

	//buffer.WriteString("Cookie:\n")
	//for _, cookie := range r.Cookies(){
	//	buffer.WriteString(fmt.Sprintf("\tCOOKIE[\"%v\"]: %v\n", cookie.Name, cookie.Value))
	//}

	buffer.WriteString("Header:\n")
	for key, value := range r.Header{
		buffer.WriteString(fmt.Sprintf("\tHeader[\"%v\"]: %v\n", key, value))
	}

	message := buffer.String()
	log.Printf(message)
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
