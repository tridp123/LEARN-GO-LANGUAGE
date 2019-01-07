package entities

import "fmt"

type Chicken struct {
	Name string
}

func (chicken Chicken) Type() string {
	return chicken.Name
}

func (chicken Chicken) Speak() {
	fmt.Println("Chicken ", chicken.Name, " vcl")
}
