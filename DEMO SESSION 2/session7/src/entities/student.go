package entities

type Student struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address Address
}

type Address struct {
	Street, Ward string
}
