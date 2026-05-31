// 🏆 Оптимальная версия (Boyer-Moore, O(1) памяти):
package main

func majorityElementII(nums []int) int {
	candidate, count := nums[0], 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate // по условию мажоритарный элемент гарантирован
}
