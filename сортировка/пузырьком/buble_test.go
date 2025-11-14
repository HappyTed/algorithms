package main

import (
	"testing"
)

var result []int

func BenchmarkBubleSorting(b *testing.B) {
	var r []int
	for i := 0; i < b.N; i++ {
		r = bubbleSort([]int{5, 7, 1, -9, 12, -22, 3, 6})
	}
	result = r
}
