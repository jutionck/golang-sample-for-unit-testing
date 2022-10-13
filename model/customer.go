package model

import "fmt"

type Customer struct {
	Id      string `json:"id"`
	Nama    string `json:"name"`
	Address string `json:"address"`
}

func (c Customer) String() {
	fmt.Println(c.Id, c.Nama, c.Address)
}

func NewCustomer(id, name, address string) Customer {
	return Customer{
		Id:      id,
		Nama:    name,
		Address: address,
	}

}
