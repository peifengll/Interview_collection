package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("get request ")
		fmt.Fprintf(w, "Hello, get")
	} else if r.Method == "POST" {
		log.Println("post request ")
		fmt.Fprintf(w, "Hello, post")

	} else {
		log.Println("other request")
		fmt.Fprintf(w, "Hello")
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
