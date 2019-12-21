package main

import "fmt"

func main() {
	/*
		https://www.acmicpc.net/problem/15489
	*/
	var r, c, w, rtn int
	rtn = 0
	_, _ = fmt.Scan(&r, &c, &w)
	var temp [][]int = generate(r + w - 1)
	for i := r - 1; i < r+w-1; i++ {
		for j := 0; j < i-r+2; j++ {
			rtn = rtn + temp[i][c+j-1]
		}
	}
	fmt.Println(rtn)
}
