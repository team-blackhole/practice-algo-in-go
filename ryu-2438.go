package main

import "fmt"

func main() {
	/*
		https://www.acmicpc.net/problem/2438
	 */
	var n int
	_, _ = fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 0; j < i ; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
