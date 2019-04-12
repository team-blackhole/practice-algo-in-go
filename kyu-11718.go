package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		https://www.acmicpc.net/problem/11718 빈 줄 불가능, 각 줄의 앞 뒤에 공백 없음
		https://www.acmicpc.net/problem/11719 빈 줄 가능, 각 줄의 앞 뒤에 공백 가능
	*/
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 100; i++ {
		c, err := reader.ReadBytes('\n')
		if err != nil || len(c) > 101 {
			break
		}
		fmt.Printf("%s", string(c))
	}
}
