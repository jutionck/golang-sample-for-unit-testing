package delivery

import (
	"fmt"

	"enigmacamp.com/golang-sample/config"
	"enigmacamp.com/golang-sample/manager"
	"enigmacamp.com/golang-sample/model"
)

type Cli struct {
	useCaseManager manager.UseCaseManager
}

func (c *Cli) Run() {
	//customerO1 := model.Customer{
	//	Id:      "1",
	//	Nama:    "Joko",
	//	Address: "Ragunan",
	//}
	customerO2 := model.Customer{
		Id:      "6",
		Nama:    "Doni",
		Address: "Pasar Rebo",
	}

	customerUseCase := c.useCaseManager.CustomerUseCase()
	customerUseCase.RegisterCustomer(customerO2)
	allCustomer, _ := customerUseCase.GetAllCustomer()

	//repo := repository.NewCustomerRepository()
	//repoKampret := repository.NewCustomerKampretRepository()
	//repo.Create(customerO1)
	//repo.Create(customerO2)
	//
	//fmt.Println(repo.RetrieveAll())
	//customerUseCase := usecase.NewCustomerUseCase(repo)
	//customerUseCase.RegisterCustomer(customerO1)
	//customerUseCase.RegisterCustomer(customerO2)
	//cust := customerUseCase.FindCustomerById("2")
	//fmt.Println(cust)
	//allCustomer := customerUseCase.GetAllCustomer()
	fmt.Println(allCustomer)
}
func Console() *Cli {
	c := config.NewConfig()

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	return &Cli{useCaseManager: usecase}
}
