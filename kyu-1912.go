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
		https://www.acmicpc.net/problem/2739
		O(n2) -> O(n)으로 풀기
	*/
	var n, rtn int
	scanner := bufio.NewReader(os.Stdin)

	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)
	subSum := make([]int, n)
	s, _ := scanner.ReadString('\n')
	str := strings.Fields(s)
	for i, s := range str {
		numbers[i], _ = strconv.Atoi(s)
	}
	subSum[0] = numbers[0]
	rtn = numbers[0]
	for i := 1; i < n; i++ {
		if subSum[i-1] > 0 {
			subSum[i] = numbers[i] + subSum[i-1]
		} else {
			subSum[i] = numbers[i]
		}
		if subSum[i] > rtn {
			rtn = subSum[i]
		}
	}
	fmt.Println(rtn)
}

//
//func main() {
//	var n ,rtn int
//	_,_ = fmt.Scan(&n)
//	numbers := make([]int,n)
//	for i:=0;i<n;i++ {
//		_,_ =fmt.Scan(&numbers[i])
//		if i!=0{
//			numbers[i] = numbers[i-1]+numbers[i]
//		}
//		if rtn< numbers[i]{
//			rtn = numbers[i]
//		}
//		for k:=0;k<i;k++{
//			if rtn < numbers[i]-numbers[k] {
//				rtn = numbers[i]-numbers[k]
//			}
//		}
//	}
//
//	fmt.Println(rtn)
//}
