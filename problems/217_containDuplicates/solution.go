package main

import "fmt"

func containsDuplicate(nums []int) bool {

	seenMap := make(map[int]bool)

	for _, number := range nums {
		if seenMap[number] {
			return true
		}

		seenMap[number] = true
	}

	return false
}

func main() {
	// Test Case 1
	nums1 := []int{1, 2, 3, 1}
	fmt.Printf("Input: %v\n", nums1)
	fmt.Printf("Output: %v\n", containsDuplicate(nums1))
	fmt.Printf("Expected: true\n\n")

	// Test Case 2
	nums2 := []int{1, 2, 3, 4}
	fmt.Printf("Input: %v\n", nums2)
	fmt.Printf("Output: %v\n", containsDuplicate(nums2))
	fmt.Printf("Expected: false\n\n")

	// Test Case 3
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Printf("Input: %v\n", nums3)
	fmt.Printf("Output: %v\n", containsDuplicate(nums3))
	fmt.Printf("Expected: true\n\n")
}
