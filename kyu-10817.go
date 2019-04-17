package main

import (
	"fmt"
	"sort"
)

func main() {
	k := make([]int, 3)
	_, _ = fmt.Scan(&k[0], &k[1], &k[2])
	sort.Ints(k)
	fmt.Print(k[1])
}
