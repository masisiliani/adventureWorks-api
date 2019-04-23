package repository

import (
	"adventureWorks-api/infra/database"
	"adventureWorks-api/model"
	"fmt"
)

//GetCustomerByID return a customer by ID
func GetCustomerByID(idCustomer int) (customer model.Customer, err error) {

	rows, err := database.DB.Query(`
							SELECT 	 
								[CustomerID]
								,ISNULL([Title], '')		AS [Title]
								,ISNULL([FirstName], '')	AS [FirstName]
								,ISNULL([MiddleName], '')	AS [MiddleName]
								,ISNULL([LastName], '')		AS [LastName]
								,ISNULL([CompanyName], '')	AS [CompanyName]
								,ISNULL([CompanyName], '')	AS [CompanyName]
								,ISNULL([Phone], '')		AS [Phone]
							FROM [SalesLT].[Customer]
							WHERE 
							CustomerID = ?;`, idCustomer)

	defer rows.Close()

	if err != nil {
		fmt.Println(err)

		return customer, err
	}

	for rows.Next() {
		err = rows.Scan(&customer.CustomerID, &customer.Title, &customer.FirstName, &customer.MiddleName,
			&customer.LastName, &customer.CompanyName, &customer.EmailAddress, &customer.Phone)

		if err != nil {
			return customer, err
		}
	}

	return customer, err
}

//GetCustomerList return a list of customers
func GetCustomerList() (err error) {

	return err
}

//UpdateCustomerById update info of the customer
func UpdateCustomerByID() (err error) {

	return err
}
