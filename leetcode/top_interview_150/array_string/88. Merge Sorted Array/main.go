package main

import "slices"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for idx, val := range nums2 {
		nums1[idx+m] = val
	}
	slices.Sort(nums1)
}

func main() {
	
}
