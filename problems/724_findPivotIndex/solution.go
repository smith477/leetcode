package main

import "fmt"

func pivotIndex(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	fmt.Printf("Sum is %d\n", sum)

	leftSum := 0
	for i, val := range nums {
		if leftSum == sum-leftSum-val {
			return i
		}
		leftSum += val
	}
	return -1
}

func main() {
	// Test Case 1
	nums1 := []int{1, 7, 3, 6, 5, 6}
	expected1 := 3
	result1 := pivotIndex(nums1)
	fmt.Printf("Test 1:\nInput: %v\nExpected: %d\nGot: %d\nPass: %v\n\n",
		nums1, expected1, result1, result1 == expected1)

	// Test Case 2
	nums2 := []int{1, 2, 3}
	expected2 := -1
	result2 := pivotIndex(nums2)
	fmt.Printf("Test 2:\nInput: %v\nExpected: %d\nGot: %d\nPass: %v\n\n",
		nums2, expected2, result2, result2 == expected2)

	// Test Case 3
	nums3 := []int{2, 1, -1}
	expected3 := 0
	result3 := pivotIndex(nums3)
	fmt.Printf("Test 3:\nInput: %v\nExpected: %d\nGot: %d\nPass: %v\n\n",
		nums3, expected3, result3, result3 == expected3)

	// Edge Case: Single element
	nums4 := []int{5}
	expected4 := 0
	result4 := pivotIndex(nums4)
	fmt.Printf("Test 4 (Single element):\nInput: %v\nExpected: %d\nGot: %d\nPass: %v\n\n",
		nums4, expected4, result4, result4 == expected4)

	// Edge Case: All zeros
	nums5 := []int{0, 0, 0}
	expected5 := 0 // Leftmost pivot
	result5 := pivotIndex(nums5)
	fmt.Printf("Test 5 (All zeros):\nInput: %v\nExpected: %d\nGot: %d\nPass: %v\n\n",
		nums5, expected5, result5, result5 == expected5)
}
