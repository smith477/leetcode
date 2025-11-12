package main

import "fmt"

func maxSubArray(nums []int) int {
	sum, maxSum := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		maxSum = max(maxSum, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
		{
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
		{
			nums:     []int{-1},
			expected: -1,
		},
		{
			nums:     []int{-2, -1},
			expected: -1,
		},
		{
			nums:     []int{-2, 1},
			expected: 1,
		},
		{
			nums:     []int{8, -19, 5, -4, 20},
			expected: 21,
		},
	}

	for i, tc := range testCases {
		result := maxSubArray(tc.nums)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("Test %d: %s (nums=%v, expected=%d, got=%d)\n",
			i+1, status, tc.nums, tc.expected, result)
	}
}
