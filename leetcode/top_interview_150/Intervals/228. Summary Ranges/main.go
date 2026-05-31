package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) < 1 {
		return result
	}

	start := 0
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			if i-1 == start {
				result = append(result, strconv.Itoa(nums[start]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			}
			start = i
			continue
		}
	}
	return result
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
