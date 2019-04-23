package model

type Customer struct {
	CustomerID   int
	Title        string
	FirstName    string
	MiddleName   string
	LastName     string
	FullName     string
	CompanyName  string
	EmailAddress string
	Phone        string
}

func (customer *Customer) FillFullName() {
	customer.FullName = customer.FirstName
	if len(customer.MiddleName) > 1 {
		customer.FullName = customer.FullName + " " + customer.MiddleName
	}

	customer.FullName = customer.FullName + " " + customer.LastName
}
