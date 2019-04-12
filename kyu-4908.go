package main

import "fmt"

func main() {
	var a, b, c float64
	_, _ = fmt.Scan(&a, &b)
	c = a / b
	fmt.Print(c)
}
