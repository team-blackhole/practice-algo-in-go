package main

import (
	"fmt"
)

func main() {
	/*
		https://www.acmicpc.net/problem/1110
	*/
	var n, l, count int
	_, _ = fmt.Scan(&n)

	l = n
	for {
		l = (l/10+l%10)%10 + (l%10)*10
		count++
		if n == l {
			break
		}

	}
	fmt.Println(count)
}
