package main

import (
	"fmt"
	"reflect"
)

// Two-pointer approach for in-place deduplication
// Time: O(n), Space: O(1)
// Use k to track position for next unique element
// Compare each element with previous, if different copy to position k and increment k
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return n
	}

	k := 1
	for i := 1; i < n; i++ {
		if nums[i-1] != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	testCases := []struct {
		name         string
		nums         []int
		expectedK    int
		expectedNums []int // first k elements should match this
	}{
		{
			name:         "Example 1",
			nums:         []int{1, 1, 2},
			expectedK:    2,
			expectedNums: []int{1, 2},
		},
		{
			name:         "Example 2",
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedK:    5,
			expectedNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:         "Single element",
			nums:         []int{1},
			expectedK:    1,
			expectedNums: []int{1},
		},
		{
			name:         "No duplicates",
			nums:         []int{1, 2, 3, 4, 5},
			expectedK:    5,
			expectedNums: []int{1, 2, 3, 4, 5},
		},
		{
			name:         "All same",
			nums:         []int{2, 2, 2, 2, 2},
			expectedK:    1,
			expectedNums: []int{2},
		},
		{
			name:         "Two elements same",
			nums:         []int{1, 1},
			expectedK:    1,
			expectedNums: []int{1},
		},
		{
			name:         "Two elements different",
			nums:         []int{1, 2},
			expectedK:    2,
			expectedNums: []int{1, 2},
		},
		{
			name:         "Negative numbers",
			nums:         []int{-3, -1, -1, 0, 0, 0, 1, 2, 2},
			expectedK:    5,
			expectedNums: []int{-3, -1, 0, 1, 2},
		},
	}

	for _, tc := range testCases {
		// Make a copy for display purposes
		inputCopy := make([]int, len(tc.nums))
		copy(inputCopy, tc.nums)

		k := removeDuplicates(tc.nums)

		// Check if k matches expected
		kCorrect := k == tc.expectedK

		// Check if first k elements match expected
		numsCorrect := reflect.DeepEqual(tc.nums[:k], tc.expectedNums)

		passed := kCorrect && numsCorrect

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Input:         %v\n", inputCopy)
			fmt.Printf("  Expected k:    %d\n", tc.expectedK)
			fmt.Printf("  Got k:         %d\n", k)
			if !numsCorrect && k > 0 && k <= len(tc.nums) {
				fmt.Printf("  Expected nums: %v\n", tc.expectedNums)
				fmt.Printf("  Got nums:      %v\n", tc.nums[:k])
			}
		}
	}
}
