package entities

import "fmt"

type Employee struct {
	Human  Human
	Id     string
	salary float32
}

func (employee Employee) ToString() string {
	return fmt.Sprintf("Name: %s\nGender: %s\nId: %s\nSalary: %f\n", employee.Human.Name, employee.Human.Gender, employee.Id, employee.salary)
}
