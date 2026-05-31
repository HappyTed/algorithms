package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return -1
	}
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left <= right {
		if nums[right] == val {
			right--
			continue
		}
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
		}
		left++
	}
	return left
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
