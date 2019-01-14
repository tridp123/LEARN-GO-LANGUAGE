package main

import (
	"config"
	"entities"
	"fmt"
	"models"
)

func main() {
	Demo5()
}

func Demo1() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Failed")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, pm := range product {
				fmt.Println(pm.ToString())
			}
		}
	}
}

/*find by id*/
func Demo2() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Failed")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err2 := productModel.Find(3)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(product.ToString())
		}
	}
}

/*add product*/
func Demo3() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Failed")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product := entities.Product{
			Name:     "abc",
			Price:    111,
			Quantity: 22,
		}
		productModel.Create(&product)
		fmt.Println(product.ToString())
	}
}

/*delete product*/
func Demo4() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Failed")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		rows, err := productModel.Delete(1)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Row Affected: ", rows)
		}
	}
}

/*update product*/
func Demo5() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Failed")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err2 := productModel.Find(3)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			product.Name = "DEF"
			product.Price = 999
			product.Quantity = 11
			rows, err3 := productModel.Update(product)
			if err3 != nil {
				fmt.Println(err3)
			} else {
				fmt.Println("Row Affected: ", rows)
			}
		}
	}
}
