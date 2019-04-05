package main

import "fmt"

func main() {
	/*
		https://www.acmicpc.net/problem/1182
	*/
	var aLength, pSum int
	_, _ = fmt.Scanln(&aLength, &pSum)

	array := make([]int, aLength)
	for i := 0; i < aLength; i++ {
		_, _ = fmt.Scan(&array[i])
	}

	partialSumArray := make([]int, aLength)
	for i := 0; i < aLength; i++ {
		if i == 0 {
			partialSumArray[i] = array[i]
		} else {
			partialSumArray[i] = array[i] + partialSumArray[i-1]
		}
	}
	count := 0
	for i := 0; i < aLength; i++ {
		for j := i; j < aLength; j++ {
			if i == j {
				if partialSumArray[i] == pSum {
					count++
				}
			} else {
				if partialSumArray[j]-partialSumArray[i] == pSum {
					count++
				}
			}
		}
	}
	fmt.Print(count)
}
