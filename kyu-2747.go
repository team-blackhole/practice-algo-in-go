package main

import "fmt"

func main() {
	/*
		https://www.acmicpc.net/problem/2747
	*/
	var n int
	_, _ = fmt.Scan(&n)
	fibo := []int{0,1}
	for i := 2; i < n+1; i++ {
		fibo = append(fibo, fibo[i-1]+fibo[i-2])
	}
	fmt.Print(fibo[n])
}

////재귀로 풀기
//func main() {
//	var n int
//	_, _ = fmt.Scan(&n)
//	fmt.Print(fibo(n))
//}
//func fibo(n int) int {
//	if (n <= 1) {
//		return n
//	} else {
//		return fibo(n-1) + fibo(n-2)
//	}
//}
