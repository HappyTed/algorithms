package main

import "fmt"

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var times, result int
	times = len(nums) / 2

	var mapped = make(map[int]int)
	for _, num := range nums {
		if _, ok := mapped[num]; ok {
			mapped[num]++
			if mapped[num] > times {
				result = num
			}
		} else {
			mapped[num] = 1
		}
	}
	return result
}

func main() {
	fmt.Println()
}
