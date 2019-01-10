package demo_json

import (
	"bufio"
	"encoding/json"
	"entities"
	"fmt"
	"os"
)

func Demo1() {
	product := entities.Product{"p01", "name 1", 2, 6}
	strjson, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(strjson))
	}
}
func Demo2() {
	str := `{"Id":"p01","Name":"name 1","Price":2,"Quantity":6}`
	var product entities.Product
	json.Unmarshal([]byte(str), &product)

	fmt.Println(product.ToString())

}

func Demo3() {
	str := `[
		{"Id":"p01","Name":"name 1","Price":2,"Quantity":16},
		{"Id":"p02","Name":"name 2","Price":3,"Quantity":26},
		{"Id":"p03","Name":"name 3","Price":4,"Quantity":36}
	]`
	var products []entities.Product

	json.Unmarshal([]byte(str), &products)

	for _, pro := range products {
		fmt.Println(pro.ToString())
	}
}

func Demo4() {
	str := `{
		"Id":"p01",
		"Name":"name 1",
		"Address":
			{
				"Street":"Street 1",
				"Ward":"Ward 1"
			}
		}`

	var student entities.Student
	err := json.Unmarshal([]byte(str), &student)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Info Student:")
		fmt.Println("Id: ", student.Id)
		fmt.Println("Name: ", student.Name)
		fmt.Println("Street: ", student.Address.Street)
		fmt.Println("Ward: ", student.Address.Ward)
	}

}

func Demo5() {
	file, err := os.Open("demo_json/invoice.json")
	// no exists
	if err != nil {
		fmt.Println(err)
	} else {
		var invoice []entities.Invoice
		scanner := bufio.NewScanner(file)
		var strjson string
		for scanner.Scan() {
			strjson = scanner.Text()
		}
		err := json.Unmarshal([]byte(strjson), &invoice)
		if err != nil {
			fmt.Println(err)
		} else {

			for _, iv := range invoice {
				fmt.Println(iv)
			}
		}

	}
	file.Close()
}
