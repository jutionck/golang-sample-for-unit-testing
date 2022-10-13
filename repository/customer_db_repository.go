package repository

import (
	"database/sql"

	"enigmacamp.com/golang-sample/model"
)

type CustomerRepository interface {
	Create(newCustomer model.Customer) error
	RetrieveAll() ([]model.Customer, error)
	FindById(id string) (model.Customer, error)
}

type customerDbRepository struct {
	db *sql.DB
}

func (c *customerDbRepository) Create(newCustomer model.Customer) error {
	insertStatement := "insert into customer values($1,$2,$3)"
	_, err := c.db.Exec(insertStatement, newCustomer.Id, newCustomer.Nama, newCustomer.Address)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerDbRepository) RetrieveAll() ([]model.Customer, error) {
	rows, err := c.db.Query("select * from customer")
	if err != nil {
		return nil, err
	}
	var customers []model.Customer
	for rows.Next() {
		var id string
		var name string
		var address string
		err = rows.Scan(&id, &name, &address)
		if err != nil {
			return nil, err
		}

		customer := model.NewCustomer(id, name, address)
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *customerDbRepository) FindById(id string) (model.Customer, error) {
	rows, err := c.db.Query("select * from customer where id=$1", id)
	if err != nil {
		return model.Customer{}, err
	}
	var customer model.Customer
	for rows.Next() {
		var id string
		var name string
		var address string
		err = rows.Scan(&id, &name, &address)
		if err != nil {
			panic(err)
		}

		customer = model.NewCustomer(id, name, address)
	}
	return customer, nil
}

// Func seperti consturcor
func NewCustomerDbRepository(db *sql.DB) CustomerRepository {
	return &customerDbRepository{
		db: db,
	}
}
