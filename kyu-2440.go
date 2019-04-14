package main

import "fmt"

func main() {
	/*
		https://www.acmicpc.net/problem/2440
	*/
	var n int
	fmt.Scan(&n)
	for k := 0; k < n; k++ {
		for i := n; i > k; i-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
