package entities

type Category struct {
	Id, Name string
	Products []Product
}

func (category Category) ToString() (result string) {
	result = "id: " + category.Id + "Name :" + category.Name + "Product :"
	for _, ca := range category.Products {
		result += ca.ToString() + "\n----------"
	}
	return
}

func (category Category) Total() (result float32) {
	for _, pr := range category.Products {
		result += pr.Total()
	}
	return
}

func (category Category) Max() (m Product) {
	m = category.Products[0]
	for _, pr := range category.Products {
		if m.Price < pr.Price {
			m = pr
		}
	}
	return
}
