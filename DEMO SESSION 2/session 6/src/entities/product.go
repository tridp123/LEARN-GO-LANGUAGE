package entities

import "fmt"

type Product struct {
	Id, Name string
	Price    float32
	Quatity  int
}

//constructor
func NewProduct(id, name string, price float32, quatity int) Product {
	return Product{id, name, price, quatity}
}

func (product Product) ToString() string {
	return fmt.Sprintf("id: %s\nName: %s\nPrice: %f\nQuatity: %d", product.Id, product.Name, product.Price, product.Quatity)
}

func (product Product) Total() float32 {
	return product.Price * float32(product.Quatity)
}
