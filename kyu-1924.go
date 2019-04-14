package main

import (
	"fmt"
)

func main() {
	/*
		https://www.acmicpc.net/problem/1924
	*/
	cal := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	week := [7]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}
	var m, d, days int = 0, 0, 0
	_, _ = fmt.Scan(&m, &d)
	for i := 0; i < m-1; i++ {
		days = days + cal[i]
	}
	fmt.Print(week[(days+d-1)%7])
}
