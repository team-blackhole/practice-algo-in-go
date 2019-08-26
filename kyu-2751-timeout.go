package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	list := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&list[i])
	}
	fmt.Println(mergeSort(list))

}

func merge(left []int, right []int) []int {
	if len(left) == 0 {
		return right
	}
	if len(right) == 0 {
		return left
	}

	var sortedList []int
	currentLeft := 0
	currentRight := 0

	for currentLeft < len(left) && currentRight < len(right) {
		if left[currentLeft] <= right[currentRight] {
			sortedList = append(sortedList, left[currentLeft])
			currentLeft += 1

			if currentLeft == len(left) {
				sortedList = append(sortedList, right[currentRight:]...)
				return sortedList
			}
		} else if left[currentLeft] > right[currentRight] {
			sortedList = append(sortedList, right[currentRight])
			currentRight += 1

			if currentRight == len(right) {
				sortedList = append(sortedList, left[currentLeft:]...)
				return sortedList
			}
		}
	}

	return sortedList
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2
	left := mergeSort(data[:mid])
	right := mergeSort(data[mid:])

	return merge(left, right)
}
