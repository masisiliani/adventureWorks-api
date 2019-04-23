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

func (customer Customer) FillFullName() {
	customer.FullName = customer.FirstName + " " + customer.MiddleName + " " + customer.LastName
}
