package main

import (
	"adventureWorks-api/infra/database"
	"adventureWorks-api/service"
	"fmt"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	customer, err := service.GetCustomerByID(1, db)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(customer.FullName)

}
