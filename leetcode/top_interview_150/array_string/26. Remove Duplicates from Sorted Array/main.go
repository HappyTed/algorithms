package main

func removeDuplicates(nums []int) int {
	var (
		slow = 0
		fast = 1
	)

	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func main() {}
