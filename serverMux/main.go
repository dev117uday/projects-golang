package main

import (
	"fmt"
	"net/http"
)

type CustomServerMux struct {
}

func (p *CustomServerMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		_, _ = fmt.Fprintf(writer, "Hello World")
		fmt.Println("\"/\" accessed")
		return
	}
	http.NotFound(writer, request)
	return
}

func main() {
	mux := &CustomServerMux{}
	fmt.Println("running at : http://localhost:8000")
	_ = http.ListenAndServe(":8000", mux)
}
