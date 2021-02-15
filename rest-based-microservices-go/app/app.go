package app

import (
	"fmt"
	"log"
	"net/http"
	"rest-based-microservices-go/domain"
	"rest-based-microservices-go/service"

	"github.com/gorilla/mux"
)

// Start : start server
func Start() {
	router := mux.NewRouter()
	var ch = CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)


	fmt.Println("Running at : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
