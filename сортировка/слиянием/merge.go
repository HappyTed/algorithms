package main

import "fmt"

func mergeSorting(lst []int) []int {
	merge := func(left, right []int) []int {
		result := make([]int, 0)
		i, j := 0, 0

		for i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				result = append(result, left[i])
				i++
			} else {
				result = append(result, right[j])
				j++
			}
		}
		result = append(result, left[i:]...)
		result = append(result, right[j:]...)

		return result
	}

	if len(lst) <= 1 {
		return lst
	}
	mid := len(lst) / 2
	left := mergeSorting(lst[:mid])
	right := mergeSorting(lst[mid:])

	return merge(left, right)
}

func main() {
	lst := []int{1, 6, 8, 9, 9, 3, -5}
	fmt.Println(lst)
	fmt.Println(mergeSorting(lst))
}
