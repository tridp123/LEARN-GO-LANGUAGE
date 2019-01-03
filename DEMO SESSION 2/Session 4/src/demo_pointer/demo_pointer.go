package demo_pointer

import (
	"fmt"
)

func Demo1() {
	var a int = 5

	fmt.Println(a)
	fmt.Println(&a)
}

func Demo2() {
	var a int = 5
	var p *int = &a
	fmt.Println(a)
	fmt.Println(&a)
	//address or value
	fmt.Println(p)
	fmt.Println(*p)

	*p = 6
	fmt.Println(a)

	q := &a
	fmt.Println(q)
	fmt.Println(*q)

	*q = 7

}
func Change1(a int) {
	a = 11
}
func Change2(a *int) {
	*a = 33
}

func Demo3() {
	a := 2
	Change1(a)
	fmt.Println(a)

	Change2(&a)
	fmt.Println(a)

}

//swap doi gia tri 2 tham so a,b (int) su dung con tro
//
func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func Demo4(p, q *int) {
	a, b := 5, 9
	p, q = &a, &b

	Swap(p, q)
	fmt.Println(a)
	fmt.Println(b)
}

func Modify1(a [3]int) {
	a[1] = 11

}

func Modify2(a *[3]int) {
	(*a)[1] = 16

}

func Demo5() {
	var a = [3]int{5, 8, 6}
	Modify1(a)
	fmt.Println(a)
	Modify2(&a)
	fmt.Println(a)
}
