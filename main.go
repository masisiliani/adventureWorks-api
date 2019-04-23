package main

import (
	"adventureWorks-api/infra/database"
	"fmt"
)

func main() {
	err := database.Connect()
	defer database.DataBaseSQL.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(database.DataBaseSQL)
}
