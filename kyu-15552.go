package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	var buf bytes.Buffer
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		str := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(str[0])
		b, _ := strconv.Atoi(str[1])
		buf.WriteString(strconv.Itoa(a + b))
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

/*
func differentSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, a, b int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	for n > 0 {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &a, &b)
		fmt.Fprintln(writer, a+b)
		n--
	}
	writer.Flush()
}
*/
