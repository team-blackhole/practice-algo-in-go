package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	score := make([]float64, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&score[i])
	}
	sort.Float64s(score)
	var sum float64
	for i := 0; i < n; i++ {
		score[i] = score[i] / score[n-1] * 100
		sum = sum + score[i]
	}
	fmt.Print(sum / float64(n))
}
