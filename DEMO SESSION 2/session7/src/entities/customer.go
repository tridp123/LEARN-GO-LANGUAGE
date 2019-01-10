package entities

type Customer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"date_of_birthday"`
}
