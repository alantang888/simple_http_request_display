package main

import (
	"net/http"
	"log"
	"fmt"
)

func requestHandler (w http.ResponseWriter, r *http.Request){
	message := fmt.Sprintf("Client from %v, request path: %v\n", r.RemoteAddr, r.URL.Path)
	log.Printf(message)
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
