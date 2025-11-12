package main

import "fmt"

func runningSum(nums []int) []int {

	result := make([]int, len(nums))

	result[0] = nums[0]

	for index := 1; index < len(nums); index++ {
		result[index] = result[index-1] + nums[index]
	}

	return result
}

func main() {
	// Test Case 1
	nums1 := []int{1, 2, 3, 4}
	expected1 := []int{1, 3, 6, 10}
	result1 := runningSum(nums1)
	fmt.Printf("Test 1:\nInput: %v\nExpected: %v\nGot: %v\nPass: %v\n\n",
		nums1, expected1, result1, equal(result1, expected1))

	// Test Case 2
	nums2 := []int{1, 1, 1, 1, 1}
	expected2 := []int{1, 2, 3, 4, 5}
	result2 := runningSum(nums2)
	fmt.Printf("Test 2:\nInput: %v\nExpected: %v\nGot: %v\nPass: %v\n\n",
		nums2, expected2, result2, equal(result2, expected2))

	// Test Case 3
	nums3 := []int{3, 1, 2, 10, 1}
	expected3 := []int{3, 4, 6, 16, 17}
	result3 := runningSum(nums3)
	fmt.Printf("Test 3:\nInput: %v\nExpected: %v\nGot: %v\nPass: %v\n\n",
		nums3, expected3, result3, equal(result3, expected3))

	// Edge Case: Single element
	nums4 := []int{5}
	expected4 := []int{5}
	result4 := runningSum(nums4)
	fmt.Printf("Test 4 (Single element):\nInput: %v\nExpected: %v\nGot: %v\nPass: %v\n\n",
		nums4, expected4, result4, equal(result4, expected4))
}

// Helper function to compare slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
