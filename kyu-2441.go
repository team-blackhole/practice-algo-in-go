package main

import "fmt"

func main() {
	/*
		https://www.acmicpc.net/problem/2441
	 */
	var n int
	_, _ = fmt.Scanln(&n)

	for i := n; i > 0; i-- {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := i; k > 0; k-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
