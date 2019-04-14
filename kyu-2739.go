package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		https://www.acmicpc.net/problem/2739
	*/
	var i int
	_, _ = fmt.Scan(&i)
	for j := 1; j < 10; j++ {
		fmt.Println(strconv.Itoa(i) + " * " + strconv.Itoa(j) + " = " + strconv.Itoa(i*j))
	}
}
