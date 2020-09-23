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
	http.ListenAndServe(":3000", nil)

}
