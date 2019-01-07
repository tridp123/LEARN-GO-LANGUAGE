package entities

import "fmt"

type Cat struct {
	Name string
}

func (cat Cat) Type() string {
	return cat.Name
}

func (cat Cat) Speak() {
	fmt.Println("Chicken ", cat.Name, " vcl")
}
