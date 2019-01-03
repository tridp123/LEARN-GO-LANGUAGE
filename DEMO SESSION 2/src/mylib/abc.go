package mylib

import (
	"fmt"
	"sort"
)

func Hello() {
	fmt.Println("tri dp")
}
func Demo6() {

	names := []string{"a", "b", "c"}
	c := append(names, "d")
	fmt.Println(c)
}
func Demo7() []int {

	names := []int{1, 2, 78, 9}
	var result []int
	for _, v := range names {
		if v > 5 {
			result = append(result, v)
		}
	}
	return result
}

func Demo8(a []int) (result1 []int, result2 []int) {
	for _, v := range a {
		if v > 0 {
			result1 = append(result1, v)
		}
		if v > 9 {
			result2 = append(result2, v)
		}
	}
	return result1, result2
}
func Demo9() {
	names := []int{1, 2, 78, 9}
	//tang
	// sort.Ints(names)
	// fmt.Println(names)
	//desc

	// sort.IntsAreSorted(names)
	fmt.Println(sort.IntsAreSorted(names))
}
