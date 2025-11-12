package main

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	start, end := -1, -2 // initialize so that if already sorted, end - start + 1 = 0
	maxSoFar := nums[0]
	minSoFar := nums[n-1]

	for i := 1; i < n; i++ {
		maxSoFar = max(maxSoFar, nums[i])
		if nums[i] < maxSoFar {
			end = i
		}
	}

	for i := n - 2; i >= 0; i-- {
		minSoFar = min(minSoFar, nums[i])
		if nums[i] > minSoFar {
			start = i
		}
	}

	return end - start + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{1}, 0},
		{[]int{1, 3, 2, 2, 2}, 4},
		{[]int{2, 1}, 2},
	}

	for i, test := range tests {
		result := findUnsortedSubarray(test.nums)
		fmt.Printf("Test %d - Expected: %d, Got: %d\n", i+1, test.output, result)
	}
}
