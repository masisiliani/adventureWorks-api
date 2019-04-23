package test

import (
	"adventureWorks-api/service"
	"fmt"
)

func TGetCustomerByID() {

	customer, err := service.GetCustomerByID(7)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(customer.FullName)

	customer, err = service.GetCustomerByID(12)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(customer.FullName)

	customer, err = service.GetCustomerByID(5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(customer.FullName)
}
