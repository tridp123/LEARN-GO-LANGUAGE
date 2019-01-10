package demo_file

import (
	"bufio"
	"entities"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Demo1() {
	file, err := os.Create("demo_file/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
	}
}

func Demo2() {
	file, err := os.Stat("demo_file/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Name: ", file.Name())
		fmt.Println("Size: ", file.Size())
		fmt.Println("Permission: ", file.Mode())
		//co so 8
		fmt.Printf("Permission: %04o", file.Mode().Perm())

	}
}

func Demo3() {
	err := os.Remove("demo_file/a.txt")
	if err != nil {
		fmt.Println(err)
	}
}

/*Read file
demo4
demo5
demo6
*/
func Demo4() {
	result, err := ioutil.ReadFile("demo_file/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}
func Demo5() {
	file, err := os.Open("demo_file/a.txt")
	// no exists
	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			//line
			line := scanner.Text()
			fmt.Println(line)
		}
	}
	file.Close()
}

func Demo6() {
	file, err := os.Open("demo_file/products.csv")
	// no exists
	if err != nil {
		fmt.Println(err)
	} else {
		var products []entities.Product
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			//line
			line := scanner.Text()
			pro := strings.Split(line, ",")
			price, _ := strconv.ParseFloat(pro[2], 64)
			quantity, _ := strconv.ParseInt(pro[3], 10, 64)
			product := entities.Product{pro[0], pro[1], price, quantity}

			fmt.Println(product.ToString())
			products = append(products, product)
		}
	}
	file.Close()
}

/*
	Write file
	demo7
	demo8
*/

func Demo7() {
	file, err := os.Create("demo_file/b.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString("Hello World\n")
	}
	file.Close()
}

func Demo8() {
	file, err := os.OpenFile("demo_file/b.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(file, "Hi tri")
	}
	file.Close()

}


