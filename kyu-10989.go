package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		https://www.acmicpc.net/problem/10989
		buffer 로 io 속도를 극복해보자.
	*/
	var n int
	_, _ = fmt.Scan(&n)
	indexList := make([]int, 100001)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		indexList[num]++
	}
	/*
	indexList := make([]int, 100001)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&temp)
		indexList[temp]++
	}
	*/

	writer := bufio.NewWriter(os.Stdout)
	for k := 1; k < 100001; k++ {
		if indexList[k] != 0 {
			for j := 0; j < indexList[k]; j++ {
				fmt.Fprintln(writer, k)
			}
		}
	}
	writer.Flush()

	/*
	for k := 1; k < 100001; k++ {
		if indexList[k] != 0 {
			for j := 0; j < indexList[k]; j++ {
			fmt.Println(k)
		}
	}
	*/

}
