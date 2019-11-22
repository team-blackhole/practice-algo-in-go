package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var line []byte

	for i := 0; i < n; i++ {
		stack := stack([]byte{})
		line, _, _ = reader.ReadLine()

		isVPS := true
		for k := 0; k < len(line); k++ {
			if line[k] == 40 {
				stack.push(line[k])
				continue
			}

			if stack.isEmpty() {
				isVPS = false
				break
			}
			stack.pop()
		}

		if isVPS && stack.isEmpty() {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

type stack []byte

func (s *stack) push(b byte) {
	*s = append(*s, b)
}

func (s *stack) pop() {
	l := len(*s)
	//n := (*s)[l-1]
	*s = (*s)[:l-1]
}

func (s stack) top() byte {
	return s[len(s)-1]
}

func (s stack) isEmpty() bool {
	if len(s) == 0 {
		return true
	} else {
		return false
	}
}
