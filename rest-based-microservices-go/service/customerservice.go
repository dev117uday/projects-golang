package service

import "rest-based-microservices-go/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer,error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer () ([]domain.Customer,error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}