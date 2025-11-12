package main

import "fmt"

func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}

		maxReach = max(maxReach, i+nums[i])

		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return maxReach >= len(nums)-1
}

func main() {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{1, 1, 1, 0}, true},
		{[]int{1, 0, 1, 0}, false},
		{[]int{2, 0, 0}, true},
	}

	for i, test := range tests {
		result := canJump(test.nums)
		status := "✓"
		if result != test.expected {
			status = "✗"
		}
		fmt.Printf("Test %d: %v → %v (expected %v) %s\n",
			i+1, test.nums, result, test.expected, status)
	}
}
