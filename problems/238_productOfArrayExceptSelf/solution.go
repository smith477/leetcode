package main

import (
	"fmt"
	"reflect"
)

// Two-pass approach: prefix and suffix products
// Time: O(n), Space: O(1) - output array doesn't count as extra space
// Pass 1: Calculate prefix products - result[i] = product of all elements left of i
// Pass 2: Calculate suffix products on-the-fly and multiply into result[i]
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func main() {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "Two elements",
			nums:     []int{2, 3},
			expected: []int{3, 2},
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "Contains zero at start",
			nums:     []int{0, 1, 2, 3},
			expected: []int{6, 0, 0, 0},
		},
		{
			name:     "Contains zero at end",
			nums:     []int{1, 2, 3, 0},
			expected: []int{0, 0, 0, 6},
		},
		{
			name:     "Two zeros",
			nums:     []int{1, 0, 2, 0, 3},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-2, -3, -4},
			expected: []int{12, 8, 6},
		},
		{
			name:     "Mix positive and negative",
			nums:     []int{2, -3, 4, -5},
			expected: []int{60, -40, 30, -24},
		},
	}

	for _, tc := range testCases {
		result := productExceptSelf(tc.nums)
		passed := reflect.DeepEqual(result, tc.expected)

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Input:    %v\n", tc.nums)
			fmt.Printf("  Expected: %v\n", tc.expected)
			fmt.Printf("  Got:      %v\n", result)
		}
	}
}
