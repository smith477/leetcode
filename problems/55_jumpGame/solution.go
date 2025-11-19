package main

import "fmt"

// Greedy approach working backwards from end
// Time: O(n), Space: O(1)
// Start from last position as goal, move goal backwards if current position can reach it
// If goal reaches index 0, the start can reach the end
func canJump(nums []int) bool {
	goal := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}

	return goal == 0
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
