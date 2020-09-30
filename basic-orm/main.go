package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitialMigration()
	handleRequest()
}

func handleRequest() {

	fmt.Println("Running at : http://localhost:8000")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", GetAllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", CreateNewUser).Methods("POST")
	myRouter.HandleFunc("/delete/{email}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/update/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", myRouter))

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
