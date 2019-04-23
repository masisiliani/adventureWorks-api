package main

import (
	"adventureWorks-api/infra/database"
	"fmt"
)

func main() {
	db, err := database.Connect()
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(db)
}
