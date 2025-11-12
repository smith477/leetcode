package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {

	if len(nums) == 1 {
		return float64(nums[0])
	}

	if len(nums) < 1 {
		return float64(0)
	}

	windowSum := 0

	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	sum := windowSum

	for i := k; i < len(nums); i++ {
		windowSum = windowSum - nums[i-k] + nums[i]
		sum = max(windowSum, sum)
	}

	return float64(sum) / float64(k)
}

func main() {
	// Test cases
	testCases := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
		{[]int{5}, 1, 5.00000},
		{[]int{0, 4, 0, 3, 2}, 1, 4.00000},
		{[]int{-1}, 1, -1.00000},
		{[]int{1, 2, 3, 4, 5}, 3, 4.00000},
		{[]int{-1, -2, -3, -4}, 2, -1.50000},
		{[]int{8, 9, 10, 11}, 4, 9.50000},
	}

	fmt.Println("Running test cases for Maximum Average Subarray I...\n")

	for i, tc := range testCases {
		result := findMaxAverage(tc.nums, tc.k)

		// Check if result is within acceptable error (10^-5)
		passed := math.Abs(result-tc.expected) < 1e-5

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("Test %d: %s\n", i+1, status)
		fmt.Printf("  Input:    nums = %v, k = %d\n", tc.nums, tc.k)
		fmt.Printf("  Expected: %.5f\n", tc.expected)
		fmt.Printf("  Got:      %.5f\n", result)
		fmt.Println()
	}
}
