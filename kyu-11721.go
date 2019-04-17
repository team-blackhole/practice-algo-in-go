package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	for i := 1; i < len(str)+1; i++ {
		fmt.Print(string(str[i-1]))
		if i%10 == 0 {
			fmt.Print("\n")
		}
	}
}
