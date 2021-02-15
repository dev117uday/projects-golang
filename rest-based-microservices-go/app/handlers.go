package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"rest-based-microservices-go/service"
)

// Customer : DS of customer
type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

// CustomerHandlers : to handle customers
type CustomerHandlers struct {
	service service.CustomerService
}

func (sch *CustomerHandlers) getAllCustomer(writer http.ResponseWriter, request *http.Request) {
	customer := []Customer{
		{Name: "Uday Yadav", City: "New Delhi", ZipCode: "110092"},
		{Name: "Uday Yadav", City: "New Delhi", ZipCode: "110092"},
		{Name: "Uday Yadav", City: "New Delhi", ZipCode: "110092"},
	}

	if request.Header.Get("Content-Type") == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customer)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(customer)
	}
}
