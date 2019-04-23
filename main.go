package main

import (
	"adventureWorks-api/infra/database"
	"adventureWorks-api/test"
	"fmt"
)

func main() {
	err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer database.DB.Close()

	test.TGetCustomerByID()

}
