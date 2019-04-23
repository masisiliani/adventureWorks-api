package service

import (
	"adventureWorks-api/model"
	"adventureWorks-api/repository"
)

func GetCustomerByID(idCustomer int) (customer model.Customer, err error) {

	customer, err = repository.GetCustomerByID(idCustomer)
	customer.FillFullName()

	if err != nil {
		return customer, err
	}

	return customer, err
}
