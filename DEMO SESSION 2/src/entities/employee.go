package entities

type Address struct {
	Street, Ward, District string
}

type FullName struct {
	FirstName, LastName string
}

type Employee struct {
	Id       string
	FullName FullName
	Address  Address
	Salary   float64
}
