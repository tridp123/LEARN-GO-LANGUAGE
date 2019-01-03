package main

import (
	"fmt"
	// "mylib"
	"entities"
	"sort"
	"strings"
)

func main() {
	fmt.Println("hello")
	// mylib.Hello()
	// mylib.Demo9()

	Demo11()
}

func Demo1() {
	var student entities.Student
	student.Id = "st01"
	student.Name = "name 01"
	student.Score = 7.8
	student.Age = 20

	fmt.Println(student)
	fmt.Println("id: ", student.Id)

}

func Demo2() {
	student := entities.Student{
		Id:    "id02",
		Name:  "trinh nhat tri",
		Age:   63,
		Score: 5.6,
	}
	fmt.Println(student)

}

func Demo3() {
	student := entities.Student{"id03", "trinh nguyen nhan", 25, 5.9}
	fmt.Println(student)

}

func Demo4() {
	var employee entities.Employee
	employee.Id = "eo1"
	employee.FullName = entities.FullName{
		FirstName: "tri",
		LastName:  "trinh",
	}
	employee.Address = entities.Address{
		Street:   "Nguyen Van Troi",
		Ward:     "Tan Thoi Hiep",
		District: "12",
	}
	employee.Salary = 96000.5

	fmt.Println("id: ", employee.Id, "\nFull Name: ", employee.FullName, "\nAddress: ", employee.Address, "\nSalary: ", employee.Salary)

}

func Demo5() {
	employee2 := entities.Employee{
		"e002",
		entities.FullName{"tri", "trinh"},
		entities.Address{"a", "b", "c"},
		6300.532}
	fmt.Println("id: ", employee2.Id, "\nFull Name: ", employee2.FullName, "\nAddress: ", employee2.Address, "\nSalary: ", employee2.Salary)

}

func ToString(student entities.Student) string {
	return fmt.Sprintf("Id:%s\nName:%s\nAge:%d\nScore:%f", student.Id, student.Name, student.Age, student.Score)
}

func Demo6() {
	student := entities.Student{
		Id:    "id02",
		Name:  "trinh nhat tri",
		Age:   63,
		Score: 5.6,
	}

	fmt.Println(ToString(student))
}
func Demo7() {
	p := &entities.Student{
		Id:    "id02",
		Name:  "trinh nhat tri",
		Age:   63,
		Score: 5.6,
	}

	fmt.Println("id: ", (*p).Id)
	fmt.Println("Name: ", (*p).Name)
}

func Change1(student entities.Student) {
	student.Name = "Nhan"
}
func Change2(student *entities.Student) {
	(*student).Name = "Nhan"
}

func Demo8() {
	student := entities.Student{
		Id:    "id02",
		Name:  "trinh nhat tri",
		Age:   63,
		Score: 5.6,
	}
	Change1(student)
	fmt.Println(ToString(student))

	Change2(&student)
	fmt.Println(ToString(student))
}

func Demo9() {
	student := [3]entities.Student{
		{"id01", "trinh nguyen nhan", 25, 5.9},
		{"id02", "trinh nhat tri", 20, 10.0},
		{"id03", "trinh thi thanh hong", 16, 9.3},
	}

	for _, v := range student {
		fmt.Println(ToString(v))
		fmt.Println("------------------------------")
	}
}

func Demo10() {
	listProduct := []entities.Product{
		{"sp01", "product 1", 100.5, 63},
		{"sp02", "product 2", 100.4, 150},
		{"sp03", "product 3", 100.7, 120},
		{"sp04", "product 4", 100.6, 100},
		{"sp05", "product 5", 100.2, 100},
	}
	fmt.Println("total Price: ", CountTotalPrice(listProduct))
	min, max := findProductMinMax(listProduct)
	fmt.Println("min: ", min)
	fmt.Println("max: ", max)

	SortProduct(listProduct[:])
	fmt.Println(listProduct)

	fmt.Println("get: ", GetProduct("pro", listProduct))

}

//a tong tien all pr float32
func CountTotalPrice(listProduct []entities.Product) (result float32) {
	result = 0
	for _, p := range listProduct {
		result += p.Price * float32(p.Quatity)

	}
	return result
}

//sp co price max, min => tra ve 2 sp
func findProductMinMax(listProduct []entities.Product) (min entities.Product, max entities.Product) {
	min = listProduct[0]
	max = listProduct[0]
	for _, p := range listProduct {
		if min.Price > p.Price {
			min = p
		}
		if max.Price < p.Price {
			max = p
		}
	}
	return min, max

}

//sort tang, giam
func SortProduct(listProduct []entities.Product) {

	sort.Slice(listProduct, func(i, j int) bool {
		return listProduct[i].Price > listProduct[j].Price
	})

}

//d search []product co Name is %key%

func GetProduct(key string, listProduct []entities.Product) (result []entities.Product) {
	for _, pro := range listProduct {
		if strings.Contains(strings.ToLower(pro.Name), strings.ToLower(key)) {
			result = append(result, pro)
		}
	}
	return result
}

func Demo11() {
	student1 := entities.Student{"id01", "trinh  nguyen nhan", 25, 5.9}
	student2 := entities.Student{"id01", "trinh nguyen nhan", 25, 5.9}
	if student1 == student2 {
		fmt.Println("student1=student2")
	} else {
		fmt.Println("student1!=student2")

	}
}
