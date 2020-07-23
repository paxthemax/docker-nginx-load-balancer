package main

import (
	"fmt"
	"log"
	"net/http"
)

// Port is the port we are serving from.
const Port = 5000

// handleHello GET /hello
func handleHello(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method, r.RequestURI)

	// Returns hello world! as a response
	fmt.Fprintln(w, "Hello world!")
}

func main() {
	// registers handleHello to GET /hello
	http.HandleFunc("/hello", handleHello)
	// starts the server on port 5000
	if err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil); err != nil {
		log.Fatalln(err)
	}
}
