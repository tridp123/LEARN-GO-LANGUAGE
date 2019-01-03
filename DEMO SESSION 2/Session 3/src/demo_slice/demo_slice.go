package demo_slice

import (
	"fmt"
	"sort"
)

func Demo1() {
	var a = [5]int{1, -4, 6, -9, 54}
	fmt.Println(a)
	a1 := a[0:2]

	fmt.Println(a1)

	a1[1] = 999
	fmt.Println(a)

	a2 := a[0:4]
	fmt.Println(a2)
	a2[1] = 555
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
}

func Demo2() {

	var a = [5]int{1, -4, 6, -9, 54}
	a1 := a[0:2]

	fmt.Println(a)

	a1 = append(a1, 123)
	fmt.Println(a1)
	fmt.Println(a)
}

func Demo3() {

	var a = [5]int{1, -4, 6, -9, 54}
	a1 := a[:2]
	fmt.Println(a1)

	a2 := a[2:]
	fmt.Println(a2)

	a3 := a[:]
	fmt.Println(a3)
}
func Demo4() {

	var a = [5]int{1, -4, 6, -9, 54}
	fmt.Println(len(a))

	a1 := a[:2]
	fmt.Println("len(a1): ", len(a1))
	fmt.Println("cap(a1): ", cap(a1))
	a2 := a[1:4]
	fmt.Println("len(a2): ", len(a2))
	fmt.Println("cap(a2): ", cap(a2))

	a3 := a[4:]
	fmt.Println("len(a3): ", len(a3))
	fmt.Println("cap(a3): ", cap(a3))

	a3 = append(a3, 888)
	fmt.Println("len(a3): ", len(a3))
	//bat tu tu slice => ket thuc mang
	fmt.Println("cap(a3): ", cap(a3))

}
func Demo5() {

	var a = make([]int, 3, 5)
	a[0] = 4
	a[1] = 6
	a[2] = -94

	a = append(a, 69, 45, 49, 85)
	fmt.Println("len(a): ", len(a))
	fmt.Println("cap(a): ", cap(a))

}
func Change1(a [3]int) {
	a[1] = 969

}
func Change2(a []int) {
	a[1] = 969
}

func Demo6() {
	var a = [3]int{1, -4, -9}
	Change1(a)
	fmt.Println(a)
	Change2(a[:])
	fmt.Println(a)

}

func Demo7() {
	var a = [5]int{1, -4, 6, -9, 54}
	fmt.Println(a)
	b := a[:]
	//sap xep b => a sap theo
	sort.Slice(b, func(i, j int) bool {
		return b[i] >= b[j]
	})
	fmt.Println(a)

}
