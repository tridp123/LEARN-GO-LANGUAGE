package demo_string_go

import (
	"fmt"
	"sort"
	"strings"
)

func Demo1() {
	s := "abcb"
	fmt.Println(len(s))
	for index := 0; index < len(s); index++ {
		fmt.Printf("\t%c", s[index])
	}
}

func Demo2() {
	s := "      abc      "
	fmt.Println(len(s))
	s = strings.TrimSpace(s)
	fmt.Println(len(s))
	for index := 0; index < len(s); index++ {
		fmt.Printf("\t%c", s[index])
	}
}
func Demo3() {
	var ch rune = 'a'
	fmt.Printf("\t%d", ch)
	fmt.Printf("\t%c\n", ch)
	s := "132abc456"

	result := strings.TrimFunc(s, func(c rune) bool {
		return c >= 48 && c <= 57
	})
	fmt.Println(result)
}
func Demo4() {

	s := "abc"
	s = strings.ToUpper(s)
	fmt.Println(s)

	s2 := "FAGFA"
	s2 = strings.ToLower(s2)
	fmt.Println(s2)

}

func Demo5() {
	s := "naruto"
	s2 := "to"

	result1 := strings.HasPrefix(s, s2)
	fmt.Println(result1)
	result2 := strings.HasPrefix(strings.ToLower(s), strings.ToLower(s2))
	fmt.Println(result2)
	result3 := strings.HasSuffix(s, s2)
	fmt.Println(result3)

	result4 := strings.Contains(s, s2)
	fmt.Println(result4)

}

func Demo6() {
	name1 := "cghsfgdsgf"
	name2 := "bahfdhsgfdsf"

	result := strings.Compare(name1, name2)
	fmt.Println(result)
}

/*
1 mang 5 pt chua 5 sp
bt:
	+ tra ve danh sach cac sp co ten keyword []
	+ sap xep cac ten tang dan, giam dan
*/
func GetProduct(key string, list []string) (result []string) {
	for _, list := range list {

		if strings.Contains(strings.ToLower(list), strings.ToLower(key)) {
			result = append(result, list)
		}
	}
	return result
}
func SortProduct(list []string) []string {
	sort.Slice(list, func(i, j int) bool {

		return strings.Compare(list[i], list[j]) < 0

	})
	return list
}
