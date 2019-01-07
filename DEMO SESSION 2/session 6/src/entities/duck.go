package entities

import "fmt"

type Duck struct {
	Name string
}

func (duck Duck) Type() string {
	return duck.Name
}

func (duck Duck) Speak() {
	fmt.Println("Chicken ", duck.Name, " vcl")
}
