package entities

import "fmt"

type Account struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Status   bool   `json:"status"`
}

func (account Account) ToString() string {
	return fmt.Sprintf("UserName: %s\nPassword: %s\nAge: %d\nStatus: %v", account.UserName, account.Password, account.Age, account.Status)
}
