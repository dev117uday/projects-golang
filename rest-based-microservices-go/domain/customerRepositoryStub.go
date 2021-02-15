package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customer := []Customer{
		{Id: "1001", Name: "Uday Yadav", City: "New Delhi", ZipCode: "110092", DateOfBirth: "2001-02-02", Status: "1"},
		{Id: "1001", Name: "Uday Yadav", City: "New Delhi", ZipCode: "110092", DateOfBirth: "2001-02-02", Status: "1"},
		{Id: "1001", Name: "Uday Yadav", City: "New Delhi", ZipCode: "110092", DateOfBirth: "2001-02-02", Status: "1"},
	}
	return CustomerRepositoryStub{customer}
}
