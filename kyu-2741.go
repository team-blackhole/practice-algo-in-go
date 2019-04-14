package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	/*
		https://www.acmicpc.net/problem/2741
	*/
	var n int
	var buffer bytes.Buffer
	_, _ = fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		buffer.WriteString(strconv.Itoa(i) + "\n")
	}
	fmt.Print(buffer.String())
}
