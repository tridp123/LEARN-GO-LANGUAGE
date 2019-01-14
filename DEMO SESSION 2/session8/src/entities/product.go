package entities

import "fmt"

type Product struct {
	Id       int64
	Name     string
	Price    float64
	Quantity int
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id: %s\nName: %s\nPrice: %f\nQuantity: %d\n--------------", product.Id, product.Name, product.Price, product.Quantity)
}
