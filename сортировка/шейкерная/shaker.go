package main

import "fmt"

func shakerSort(list []int) {

	left := 0
	right := len(list) - 1

	for {
		if left > right {
			break
		}

		for i := left; i < right; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
		left += 1

		for j := right; j >= left; j-- {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
		right -= 1
	}

}

func main() {
	lst := []int{1, 6, 8, 9, 9, 3, -5}
	fmt.Println(lst)
	shakerSort(lst)
	fmt.Println(lst)
}
