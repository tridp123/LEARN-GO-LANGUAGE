package demo_map

import "fmt"

func Demo1() {
	student := make(map[string]string)

	student["id"] = "st01"
	student["name"] = "name 1"
	student["address"] = "address 1"

	fmt.Println(student)

	for key, value := range student {
		fmt.Println(key, " : ", value)
	}
}

func Demo2() {
	student := map[string]string{
		"id":      "id1",
		"name":    "name 1",
		"address": "address 1",
	}

	// student["id"] = "st01"
	// student["name"] = "name 1"
	// student["address"] = "address 1"
	student["phone"] = "12346"
	fmt.Println(student)

	for key, value := range student {
		fmt.Println(key, " : ", value)
	}
}

func Demo3() {
	student := map[string]string{
		"id":      "id1",
		"name":    "name 1",
		"address": "address 1",
	}

	// student["id"] = "st01"
	// student["name"] = "name 1"
	// student["address"] = "address 1"
	student["phone"] = "12346"
	fmt.Println(student)

	for key, value := range student {
		fmt.Println(key, " : ", value)
	}

	value, ok := student["name"]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("error")
	}

	delete(student, "address")
	fmt.Println(student)
}

func Display(student map[string]string) {
	for key, value := range student {
		fmt.Println(key, " : ", value)
	}
}

func Demo4() {
	student := map[string]string{
		"id":      "id1",
		"name":    "name 1",
		"address": "address 1",
	}
	Display(student)
}

func Demo5() {

	categories := map[string][]string{
		"category 1": []string{"product 1", "product 2"},
		"category 2": []string{"product 3", "product 4", "product 5", "product 6"},
		"category 3": []string{"product 7", "product 8"},
	}

	fmt.Println(categories)

	categories["category 4"] = []string{"product 9", "product 10"}
	fmt.Println(categories)
	categories["category 3"] = append(categories["category 3"], "product 6.1", "product 6.2", "product 6.3")
	fmt.Println(categories)

	for key, value := range categories {
		fmt.Println(key)
		for _, name := range value {
			fmt.Println(name)
		}
	}

}
