package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	http.HandleFunc("/", helloWorld)
	fmt.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from web server")
}
