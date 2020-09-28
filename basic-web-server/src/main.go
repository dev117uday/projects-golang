package main

import (
	"fmt"
	"html"
	"net/http"
	"projects-golang/basic-web-server/lib"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		urlPathElements := strings.Split(request.URL.Path, "/")
		if urlPathElements[1] == "roman" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				fmt.Println("Number accessed : ", number)
				writer.WriteHeader(http.StatusNotFound)
				_, _ = writer.Write([]byte("404 - Not Found\n"))
			} else {
				fmt.Println("Number accessed : ", number)
				fmt.Println("Numeral returned : ", lib.Numerals[number])
				_, _ = fmt.Fprintf(writer, "%q\n", html.EscapeString(lib.Numerals[number]))
			}
		} else {
			writer.WriteHeader(http.StatusNotFound)
			_, _ = writer.Write([]byte("404 - Bad Request\n"))
		}
	})
	server := &http.Server{
		Addr: ":8000",
	}
	fmt.Println("running at : http://localhost:8000")
	_ = server.ListenAndServe()
}
