package usecase

import (
	"enigmacamp.com/golang-sample/model"
	"enigmacamp.com/golang-sample/repository"
)

type CustomerUsecase interface {
	RegisterCustomer(customer model.Customer) error
	FindCustomerById(id string) (model.Customer, error)
	GetAllCustomer() ([]model.Customer, error)
}
type customerUsecase struct {
	repo repository.CustomerRepository
}

func (c *customerUsecase) RegisterCustomer(customer model.Customer) error {
	return c.repo.Create(customer)
}

func (c *customerUsecase) FindCustomerById(id string) (model.Customer, error) {
	return c.repo.FindById(id)
}

func (c *customerUsecase) GetAllCustomer() ([]model.Customer, error) {
	return c.repo.RetrieveAll()
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{repo: repo}
}
