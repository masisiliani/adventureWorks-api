package repository

import (
	"adventureWorks-api/model"
	"database/sql"
	"fmt"
)

//GetCustomerByID return a customer by ID
func GetCustomerByID(idCustomer int, db *sql.DB) (customer model.Customer, err error) {

	fmt.Println(db)
	rows, err := db.Query(`
							SELECT 	 
								[CustomerID]
								,[Title]
								,[FirstName]
								,[MiddleName]
								,[LastName]
								,[CompanyName]
								,[EmailAddress]
								,[Phone]
							FROM [SalesLT].[Customer]
							WHERE 
							CustomerID = ?`, idCustomer)

	defer rows.Close()

	if err != nil {
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
