package service

import (
	"adventureWorks-api/model"
	"adventureWorks-api/repository"
	"database/sql"
)

func GetCustomerByID(idCustomer int, db *sql.DB) (customer model.Customer, err error) {
	customer, err = repository.GetCustomerByID(idCustomer, db)
	customer.FillFullName()

	if err != nil {
		return customer, err
	}

	return customer, err
}
