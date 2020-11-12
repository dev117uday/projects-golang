package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article struct : used to store articles
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//Articles : a list to store struct Article
type Articles []Article

func main() {

	fmt.Println("running at : http://localhost:8000")
	handleRequest()

}

func handleRequest() {

	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homePage)
	myrouter.HandleFunc("/article", getArticles).Methods("GET")
	myrouter.HandleFunc("/article", replyArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myrouter))

}

func getArticles(writer http.ResponseWriter, request *http.Request) {

	articles := Articles{
		Article{
			Title:   "Trump becomes president again",
			Desc:    "Chine Helped him",
			Content: "Trump is back with all bakwas",
		},
		Article{
			Title:   "title 2",
			Desc:    "this is just a description",
			Content: "there is nothing to write in th article",
		},
	}
	fmt.Println("/article endpoint hit")
	json.NewEncoder(writer).Encode(articles)

}

func homePage(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("/ endpoint hit")
	fmt.Fprintf(writer, "Homepage endpoint hit")

}

func replyArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "reply article POST request")
}
