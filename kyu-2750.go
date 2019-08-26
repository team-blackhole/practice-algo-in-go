package main

import "fmt"

func main() {
	/*
		https://www.acmicpc.net/problem/2750
	*/
	var n int
	_, _ = fmt.Scan(&n)
	list := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&list[i])
	}

	////bubble sort
	//for j := n; j > 0; j-- {
	//	for k := 0; k < j-1; k++ {
	//		if list[k] > list[k+1] {
	//			list[k], list[k+1] = list[k+1], list[k]
	//		}
	//	}
	//}
	//
	////insertion sort
	//for j := 1; j < n; j++ {
	//	for k := 0; k <j; k++ s{
	//		if list[k]>list[j] {
	//			list[j] , list[k] = list[k], list[j]
	//		}
	//	}
	//}

	//selection sort
	for j := 0; j < n; j++ {
		min := j
		for i := j + 1; i < n; i++ {
			if list[min] >list[i] {
				min = i
			}
		}
		list[j] ,list[min] = list[min], list[j]
	}

	for m := 0; m < n; m++ {
		fmt.Println(list[m])
	}
}
