package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	//buffer := bytes.Buffer{}
	buffer := strings.Builder{}
	buffer.WriteString(fmt.Sprintf("Client from %v, request path: %v\n", r.RemoteAddr, r.URL.Path))
	buffer.WriteString("==================================================================\n\n")

	cookies := r.Cookies()
	if len(cookies) > 0 {
		buffer.WriteString("Cookie:\n")
		for _, cookie := range cookies {
			buffer.WriteString(fmt.Sprintf("\tCOOKIE[\"%v\"]: %v\n", cookie.Name, cookie.Value))
		}
		buffer.WriteString("==================================================================\n\n")
	}

	headers := r.Header
	if len(headers) > 0 {
		buffer.WriteString("Header:\n")
		for key, value := range headers {
			buffer.WriteString(fmt.Sprintf("\tHeader[\"%v\"]: %v\n", key, value))
		}
		buffer.WriteString("==================================================================\n\n")
	}

	message := buffer.String()
	log.Printf(message)
	fmt.Fprintf(w, message)
}

func main() {
	listenPortStr := os.Getenv("SERVER_PORT")
	listenPort, err := strconv.Atoi(listenPortStr)
	if err != nil {
		listenPort = 8080
	}
	fmt.Printf("Will try to listen on http://0.0.0.0:%[1]d. You may try on http://127.0.0.1:%[1]d", listenPort)
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil))
}
