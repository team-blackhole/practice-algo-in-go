package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, x int
	var str string
	_, _ = fmt.Scanln(&n, &x)
	ar := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&ar[i])
	}
	for i := 0; i < n; i++ {
		if ar[i] < x {
			str = str + strconv.Itoa(ar[i]) + " "
		}
	}
	fmt.Print(str)
}
