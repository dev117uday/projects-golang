package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndPoint)
}

func main() {

	fmt.Println("Go web sockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ws EndPoint")
}
