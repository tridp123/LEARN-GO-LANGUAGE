package entities

import "fmt"

type Human struct {
	Name, Gender string
}

func (human Human) ToString() string {
	return fmt.Sprintf("Name: %s\nGender: %s", human.Name, human.Gender)
}

func (human Human) Eat() string {
	return fmt.Sprintf("eating..")
}
func (human Human) Sleep() string {
	return fmt.Sprintf("sleeping..")
}

func (human Human) Move() string {
	return fmt.Sprintf("moving..")
}
