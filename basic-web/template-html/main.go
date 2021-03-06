package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	fmt.Println("running at : http://localhost:8080")
	http.ListenAndServe(":8000", nil)
}
