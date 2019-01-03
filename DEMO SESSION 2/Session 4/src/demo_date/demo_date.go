package demo_date

import (
	"fmt"
	"time"
)

func Demo1() {
	today := time.Now()
	fmt.Println("today: ", today)
	fmt.Println("Year: ", today.Year())
	month := today.Month()
	fmt.Println(month)
	//get number
	fmt.Println(int(month))

	fmt.Println("Day: ", today.Day())

}

func Demo2() {
	today := time.Now()
	fmt.Println(today.Format("02/01/2006 15:04:05"))

}

func Demo3() {
	s := "17/11/1996"

	d, error := time.Parse("02/01/2006", s)

	if error != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(d.Format("02/01/2006"))
	}
}

func Demo4() {
	today := time.Now()

	time1 := today.AddDate(0, 0, 2)

	fmt.Println("time 1: ", time1.Format("02/01/2006"))

	time2 := today.Add(3 * 24 * time.Hour)
	fmt.Println("time 2: ", time2.Format("02/01/2006"))
}

/*
cho 1 ngay sinh kieu chuoi.
1. kiem tra co phai sinh nhat
2. tinh tuoi dua vao ngay sinh
*/

func CheckBirthDay(date string) {
	d, error := time.Parse("02/01/2006", date)
	if error != nil {
		fmt.Println("Error")
	} else {
		today := time.Now()
		if d.Month() == today.Month() && d.Year() == today.Year() && d.Day() == today.Year() {
			fmt.Println("is birth day")
		} else {
			fmt.Println("is not birth day")
		}
	}
}

/*17/11/1996
27/12/2018
*/
// func CountOld(date string)  int{
// 	d, error := time.Parse("02/01/2006", date)
// 	today := time.Now()
// 	old := today.Year()-d.Year()
// 	if error != nil {
// 		fmt.Println("Error")
// 	} else {
// 		if today.Month()< d.Month() || (today.Month()== d.Month() && today.Day() < d.Day()){
// 			old = old -1
// 	}
// 	return old
// }
