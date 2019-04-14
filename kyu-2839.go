package main

import "fmt"

func main() {
	/*
		https://www.acmicpc.net/problem/2839
		좀 더 쉽게 푸는 방법이 있는 것 같지만...
	*/
	var a int
	_, _ = fmt.Scan(&a)

	if a%5 == 0 {
		fmt.Print(a / 5)
	} else if (a-5) > 0 && (a-5)%3 == 0 {
		fmt.Print(div(a-5) + 1)
	} else if (a-10) > 0 && (a-10)%3 == 0 {
		fmt.Print(div(a-10) + 2)
	} else if a%3 == 0 {
		fmt.Print(div(a))
	} else {
		fmt.Print(-1)
	}
}

func div(sub int) int {
	return sub/15*3 + sub%15/3
}
