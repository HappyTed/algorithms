package main

import (
	"fmt"
)

type Collection interface {
	~string | ~int | ~float32 | ~float64
}

func bubbleSort[T Collection](list []T) []T {
	n := len(list)

	if n < 2 {
		return list
	}

	for {
		swapped := false
		for i := 0; i < n-1; i++ {
			if list[i] > list[i+1] {
				swapped = true
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
		if !swapped {
			break
		}
	}
	return list
}

func main() {
	lst := []int{1, 6, 8, 9, 9, 3, -5}
	sorted := bubbleSort(lst)
	fmt.Println(sorted)
}
