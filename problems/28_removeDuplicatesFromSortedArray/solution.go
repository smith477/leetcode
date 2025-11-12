package main

import (
	"fmt"
	"reflect"
)

// Two pointer
// Time complexity O(n), Space complexity O(1)
// Pointer k tracks where to write next distinct (non-val) element
// Iterate through array with i
// When nums[i] != val, write it at position k and increment k
// When nums[i] == val, skip it (only i increments, k stays same)
// This way k only advances when writing non-val elements
// Return k as the count of non-val elements
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1 // Position for next unique element

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] { // Found a new unique element
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	// Example 1
	fmt.Println("=== Example 1 ===")
	nums1 := []int{1, 1, 2}
	expected1 := []int{1, 2}
	fmt.Printf("Before: %v\n", nums1)
	k1 := removeDuplicates(nums1)
	fmt.Printf("After:  k=%d, nums=%v (first %d elements)\n", k1, nums1[:k1], k1)
	fmt.Printf("Expected: k=%d, nums=%v\n", len(expected1), expected1)
	fmt.Printf("Match: %v\n\n", k1 == len(expected1) && reflect.DeepEqual(nums1[:k1], expected1))

	// Example 2
	fmt.Println("=== Example 2 ===")
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected2 := []int{0, 1, 2, 3, 4}
	fmt.Printf("Before: %v\n", nums2)
	k2 := removeDuplicates(nums2)
	fmt.Printf("After:  k=%d, nums=%v (first %d elements)\n", k2, nums2[:k2], k2)
	fmt.Printf("Expected: k=%d, nums=%v\n", len(expected2), expected2)
	fmt.Printf("Match: %v\n\n", k2 == len(expected2) && reflect.DeepEqual(nums2[:k2], expected2))

	// Test 3: No duplicates
	fmt.Println("=== Test 3: No Duplicates ===")
	nums3 := []int{1, 2, 3, 4, 5}
	expected3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before: %v\n", nums3)
	k3 := removeDuplicates(nums3)
	fmt.Printf("After:  k=%d, nums=%v\n", k3, nums3[:k3])
	fmt.Printf("Expected: k=%d, nums=%v\n", len(expected3), expected3)
	fmt.Printf("Match: %v\n\n", k3 == len(expected3) && reflect.DeepEqual(nums3[:k3], expected3))

	// Test 4: All same
	fmt.Println("=== Test 4: All Same ===")
	nums4 := []int{1, 1, 1, 1, 1}
	expected4 := []int{1}
	fmt.Printf("Before: %v\n", nums4)
	k4 := removeDuplicates(nums4)
	fmt.Printf("After:  k=%d, nums=%v\n", k4, nums4[:k4])
	fmt.Printf("Expected: k=%d, nums=%v\n", len(expected4), expected4)
	fmt.Printf("Match: %v\n\n", k4 == len(expected4) && reflect.DeepEqual(nums4[:k4], expected4))

	// Test 5: Single element
	fmt.Println("=== Test 5: Single Element ===")
	nums5 := []int{1}
	expected5 := []int{1}
	fmt.Printf("Before: %v\n", nums5)
	k5 := removeDuplicates(nums5)
	fmt.Printf("After:  k=%d, nums=%v\n", k5, nums5[:k5])
	fmt.Printf("Expected: k=%d, nums=%v\n", len(expected5), expected5)
	fmt.Printf("Match: %v\n\n", k5 == len(expected5) && reflect.DeepEqual(nums5[:k5], expected5))

	// Test 6: Two elements same
	fmt.Println("=== Test 6: Two Elements Same ===")
	nums6 := []int{1, 1}
	expected6 := []int{1}
	fmt.Printf("Before: %v\n", nums6)
	k6 := removeDuplicates(nums6)
	fmt.Printf("After:  k=%d, nums=%v\n", k6, nums6[:k6])
	fmt.Printf("Expected: k=%d, nums=%v\n", len(expected6), expected6)
	fmt.Printf("Match: %v\n\n", k6 == len(expected6) && reflect.DeepEqual(nums6[:k6], expected6))

	// Test 7: Two elements different
	fmt.Println("=== Test 7: Two Elements Different ===")
	nums7 := []int{1, 2}
	expected7 := []int{1, 2}
	fmt.Printf("Before: %v\n", nums7)
	k7 := removeDuplicates(nums7)
	fmt.Printf("After:  k=%d, nums=%v\n", k7, nums7[:k7])
	fmt.Printf("Expected: k=%d, nums=%v\n", len(expected7), expected7)
	fmt.Printf("Match: %v\n\n", k7 == len(expected7) && reflect.DeepEqual(nums7[:k7], expected7))

	// Test 8: Negative numbers
	fmt.Println("=== Test 8: Negative Numbers ===")
	nums8 := []int{-3, -3, -2, -1, -1, 0, 0, 0, 0, 0}
	expected8 := []int{-3, -2, -1, 0}
	fmt.Printf("Before: %v\n", nums8)
	k8 := removeDuplicates(nums8)
	fmt.Printf("After:  k=%d, nums=%v\n", k8, nums8[:k8])
	fmt.Printf("Expected: k=%d, nums=%v\n", len(expected8), expected8)
	fmt.Printf("Match: %v\n\n", k8 == len(expected8) && reflect.DeepEqual(nums8[:k8], expected8))
}
