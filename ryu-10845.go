package main

import "fmt"

func main() {
	/*
	https://www.acmicpc.net/problem/10845

	push X: 정수 X를 큐에 넣는 연산이다.
	pop: 큐에서 가장 앞에 있는 정수를 빼고, 그 수를 출력한다. 만약 큐에 들어있는 정수가 없는 경우에는 -1을 출력한다.
	size: 큐에 들어있는 정수의 개수를 출력한다.
	empty: 큐가 비어있으면 1, 아니면 0을 출력한다.
	front: 큐의 가장 앞에 있는 정수를 출력한다. 만약 큐에 들어있는 정수가 없는 경우에는 -1을 출력한다.
	back: 큐의 가장 뒤에 있는 정수를 출력한다. 만약 큐에 들어있는 정수가 없는 경우에는 -1을 출력한다.
	 */
	var nOrder int
	_, _ = fmt.Scanln(&nOrder)

	var queue []int

	for i := 0; i < nOrder; i++ {
		var order string
		var x int
		_,_ = fmt.Scanln(&order, &x)

		switch order {
		case "push":
			queue = append(queue, x)
		case "pop":
			if len(queue) == 0 {
				fmt.Println("-1")
			} else {
				fmt.Println(queue[0])
				queue = queue[1:]
			}
		case "size":
			fmt.Println(len(queue))
		case "empty":
			if len(queue) == 0 {
				fmt.Println("1")
			} else {
				fmt.Println("0")
			}
		case "front":
			if len(queue) == 0 {
				fmt.Println("-1")
			} else {
				fmt.Println(queue[0])
			}
		case "back":
			if len(queue) == 0 {
				fmt.Println("-1")
			} else {
				fmt.Println(queue[len(queue)-1])
			}
		}
	}
}
