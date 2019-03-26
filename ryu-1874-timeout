package main

import (
	"bytes"
	"fmt"
)

func main() {
	/*
		https://www.acmicpc.net/problem/1874
	*/
	var nOrder int
	_, _ = fmt.Scanln(&nOrder)

	var stack [100000]int
	lenStack := 0

	stackNext := 1

	var buf bytes.Buffer
	var x int

	for i := 0; i < nOrder; i++ {
		_, _ = fmt.Scan(&x)

		for ; ; stackNext++ {
			if lenStack > 0 {
				top := stack[lenStack-1]
				if x == top {
					lenStack--
					_, _ = buf.WriteString("-\n")
					break
				}
				if x < top {
					fmt.Println("NO")
					return
				}
			}

			if stackNext > nOrder {
				fmt.Println("NO")
				return
			}

			stack[lenStack] = stackNext
			lenStack++
			_, _ = buf.WriteString("+\n")
		}
	}

	if lenStack > 0 {
		fmt.Println("NO")
		return
	}
	fmt.Print(buf.String())
}
