package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)

	for index, num := range nums {

		if lastIndex, found := indexMap[num]; found {
			if index-lastIndex <= k {
				return true
			}
		}

		indexMap[num] = index
	}
	return false
}

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{1, 2, 3, 4, 5}, 2, false},
		{[]int{1, 1}, 0, false},
		{[]int{1, 1}, 1, true},
		{[]int{99, 99}, 2, true},
		{[]int{1, 2, 1}, 1, false},
		{[]int{1, 2, 1}, 2, true},
	}

	fmt.Println("Running test cases for Contains Duplicate II...\n")

	for i, tc := range testCases {
		result := containsNearbyDuplicate(tc.nums, tc.k)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("Test %d: %s\n", i+1, status)
		fmt.Printf("  Input:    nums = %v, k = %d\n", tc.nums, tc.k)
		fmt.Printf("  Expected: %v\n", tc.expected)
		fmt.Printf("  Got:      %v\n", result)
		fmt.Println()
	}
}
