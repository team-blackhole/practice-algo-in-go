package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
	https://www.acmicpc.net/problem/10828

	push X: 정수 X를 스택에 넣는 연산이다.
	pop: 스택에서 가장 위에 있는 정수를 빼고, 그 수를 출력한다. 만약 스택에 들어있는 정수가 없는 경우에는 -1을 출력한다.
	size: 스택에 들어있는 정수의 개수를 출력한다.
	empty: 스택이 비어있으면 1, 아니면 0을 출력한다.
	top: 스택의 가장 위에 있는 정수를 출력한다. 만약 스택에 들어있는 정수가 없는 경우에는 -1을 출력한다.
	 */
	var nOrder int
	_, _ = fmt.Scanln(&nOrder)

	var stack []int

	reader := bufio.NewReader(os.Stdin)

	var order string
	var x int
	var line []byte

	for i := 0; i < nOrder; i++ {
		line, _, _ = reader.ReadLine()

		parts := strings.Split(string(line), " ")
		order = parts[0]
		if len(parts) > 1 {
			x, _ = strconv.Atoi(parts[1])
		}

		switch order {
		case "push":
			stack = append(stack, x)
		case "pop":
			if len(stack) == 0 {
				fmt.Println("-1")
			} else {
				fmt.Println(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		case "size":
			fmt.Println(len(stack))
		case "empty":
			if len(stack) == 0 {
				fmt.Println("1")
			} else {
				fmt.Println("0")
			}
		case "top":
			if len(stack) == 0 {
				fmt.Println("-1")
			} else {
				fmt.Println(stack[len(stack)-1])
			}
		}
	}
}
