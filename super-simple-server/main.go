package main

import (
	"fmt"
	"net/http"
)

func indexhandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Print("/ is accessed")
	fmt.Fprintf(writer, "Hello World")
}

func main() {

	http.HandleFunc("/", indexhandler)
	fmt.Println("running at : http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
