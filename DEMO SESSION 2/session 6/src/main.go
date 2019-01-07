package main

import (
	"entities"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("************")
	Demo8()
}

func Demo1() {
	student := entities.Student{"st01", "name 1", 20}
	fmt.Println(student.ToString())
}

func Demo2() {
	product1 := entities.Product{"st01", "product name 1", 6.9, 20}
	product2 := entities.Product{"st02", "product name 2", 3.6, 26}
	fmt.Println(product1.ToString())
	fmt.Println(product2.Total())
}

func Demo3() {
	student := entities.NewStudent("id1", "tri", 25)
	fmt.Println(student.ToString())
	fmt.Println("-------------------")
	student.SetName("nhan")
	(&student).SetAge(15)
	fmt.Println(student.ToString())
}

func Demo4() {
	product := entities.NewProduct("st01", "product name 1", 6.9, 20)
	fmt.Println(product.ToString())

}
func Demo5() {
	category1 := entities.Category{"01", "ca01", []entities.Product{
		{"st01", "product name 1", 6.9, 20},
		{"st02", "product name 2", 2.9, 40},
	}}
	category2 := entities.Category{"02", "ca02", []entities.Product{
		{"st01", "product name 1", 6.9, 20},
		{"st02", "product name 2", 2.9, 40},
	}}
	fmt.Println(category1)
	fmt.Println(category2)

	category := []entities.Category{category1, category2}
	fmt.Println("Total: ", TotalPriceListCategory(category))
	fmt.Println("Max: ", category1.Max())
	fmt.Println("Max of List cate: " + getProductMaxAListCategory(category).ToString())

}

/*
1. tong tien tat ca san pham co trong cac danh muc
2. liet ke thong tin san pham co price cao nhat trong 1 danh muc
3. liet ke thong tin san pham co price cao nhat trong ds danh muc
*/
func TotalPriceListCategory(listCategory []entities.Category) (result float32) {
	for _, ca := range listCategory {
		result += ca.Total()
	}
	return
}

func getProductMaxAListCategory(listCategory []entities.Category) (m entities.Product) {
	m = listCategory[0].Max()
	for _, ca := range listCategory {
		if ca.Max().Price > m.Price {
			m = ca.Max()
		}
	}
	return
}

// func Demo6() {
// 	employee := entities.Employee{
// 		entities.Human{"nguyen can", "male"}, "id1", 6.9}
// 	fmt.Println(employee.ToString())
// }

func Demo7() {
	var human entities.Human
	var ihuman entities.IHuman = human

	fmt.Println(ihuman.Eat())
	fmt.Println(ihuman.Sleep())

	var ihuman2 entities.IHuman2 = human
	ihuman2.Move()

}

func Demo8() {
	animal := []entities.Animal{
		entities.Duck{"Duck 1"},
		entities.Cat{"Cat 1"},
		entities.Chicken{"Chicken 1"},
		entities.Duck{"Duck 12"},
		entities.Cat{"Cat 12"},
		entities.Chicken{"Chicken 12"},
	}
	for _, an := range animal {
		an.Speak()
	}
}
