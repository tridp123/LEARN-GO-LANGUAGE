package main

import (
	"demo_string_go"
	"fmt"
)

func main() {
	fmt.Println("Hello session3")
	var listProduct = []string{
		"product 1",
		"anh khoa 2",
		"kim dung 1",
		"kim dung 2",
	}
	result := demo_string_go.GetProduct("kim", listProduct)
	fmt.Println(result)
	fmt.Println(demo_string_go.SortProduct(listProduct))
}
