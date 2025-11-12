package main

import "fmt"

// Fibonacci pattern - at each house decide: rob it or skip it
// Time complexity O(n), Space complexity O(1)
// Track max money from previous two positions
// Base: house 0 = nums[0], house 1 = max(nums[0], nums[1])
// Recurrence: rob[i] = max(rob[i-2] + nums[i], rob[i-1])
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	robPrevHouse := nums[0]
	robCurrentHouse := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		robThisHouse := max(robPrevHouse+nums[i], robCurrentHouse)
		robPrevHouse = robCurrentHouse
		robCurrentHouse = robThisHouse
	}

	return robCurrentHouse
}

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			nums:     []int{2, 7, 9, 3, 1},
			expected: 12,
		},
		{
			nums:     []int{2, 1, 1, 2},
			expected: 4,
		},
		{
			nums:     []int{5, 3, 4, 11, 2},
			expected: 16,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
		{
			nums:     []int{2, 1},
			expected: 2,
		},
	}

	for i, tc := range testCases {
		result := rob(tc.nums)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("Test %d: %s (nums=%v, expected=%d, got=%d)\n",
			i+1, status, tc.nums, tc.expected, result)
	}
}
