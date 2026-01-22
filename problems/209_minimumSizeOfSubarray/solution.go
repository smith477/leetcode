package main

import (
	"fmt"
)

// Sliding window approach
// Time: O(n), Space: O(1)
// Expand window by adding elements, shrink from left when sum >= target
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minLen := len(nums) + 1
	windowSum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		windowSum += nums[right]

		for windowSum >= target {
			windowSize := right - left + 1
			if windowSize < minLen {
				minLen = windowSize
			}

			windowSum -= nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}

func main() {
	// Test case 1
	target1 := 7
	nums1 := []int{2, 3, 1, 2, 4, 3}
	fmt.Printf("Input: target = %d, nums = %v\n", target1, nums1)
	fmt.Printf("Output: %d\n", minSubArrayLen(target1, nums1))
	fmt.Printf("Expected: 2\n\n")

	// Test case 2
	target2 := 4
	nums2 := []int{1, 4, 4}
	fmt.Printf("Input: target = %d, nums = %v\n", target2, nums2)
	fmt.Printf("Output: %d\n", minSubArrayLen(target2, nums2))
	fmt.Printf("Expected: 1\n\n")

	// Test case 3
	target3 := 11
	nums3 := []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Printf("Input: target = %d, nums = %v\n", target3, nums3)
	fmt.Printf("Output: %d\n", minSubArrayLen(target3, nums3))
	fmt.Printf("Expected: 0\n\n")
}
