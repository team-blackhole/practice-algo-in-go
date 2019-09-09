package main

import (
	"fmt"
	"io"
)

func main() {
	/*
		https://www.acmicpc.net/problem/10951
	*/
	var a, b int

	for {
		_, err := fmt.Scanln(&a, &b)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(a + b)
	}
}
