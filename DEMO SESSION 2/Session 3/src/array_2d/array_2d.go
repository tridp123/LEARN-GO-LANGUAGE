package array_2d

import "fmt"

func Demo1() {
	var a [2][3]int
	a[0][0] = 1
	a[0][1] = 5
	a[0][2] = -9

	a[1][0] = -69
	a[1][1] = 85
	a[1][2] = 99

	fmt.Println(a)

	fmt.Println(len(a))

	for index := 0; index < len(a); index++ {
		for index2 := 0; index2 < len(a); index++ {
			fmt.Println("\t", a[index][index2])
		}
		fmt.Println()
	}
}
func Demo2() {
	a := [2][3]int{
		{4, 6, -9},
		{5, -8, -3},
	}
	fmt.Println(a)

	for i, row := range a {
		fmt.Println("row", i)
		fmt.Println("\t", row)
		for j, v := range row {
			fmt.Println("j: ", j, "var: ", v)
		}
	}
}
