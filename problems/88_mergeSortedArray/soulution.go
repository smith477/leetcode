package main

import (
	"fmt"
	"reflect"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	totalCount := m + n - 1
	first := m - 1
	second := n - 1

	for second >= 0 {
		if first >= 0 && nums1[first] > nums2[second] {
			nums1[totalCount] = nums1[first]
			first--
		} else {
			nums1[totalCount] = nums2[second]
			second--
		}
		totalCount--
	}
}

func main() {
	// Example 1
	fmt.Println("=== Example 1 ===")
	nums1_1 := []int{1, 2, 3, 0, 0, 0}
	nums2_1 := []int{2, 5, 6}
	m1, n1 := 3, 3
	fmt.Printf("Before: nums1=%v, nums2=%v, m=%d, n=%d\n", nums1_1, nums2_1, m1, n1)
	merge(nums1_1, m1, nums2_1, n1)
	expected1 := []int{1, 2, 2, 3, 5, 6}
	fmt.Printf("After:  nums1=%v\n", nums1_1)
	fmt.Printf("Expected: %v\n", expected1)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums1_1, expected1))

	// Example 2
	fmt.Println("=== Example 2 ===")
	nums1_2 := []int{1}
	nums2_2 := []int{}
	m2, n2 := 1, 0
	fmt.Printf("Before: nums1=%v, nums2=%v, m=%d, n=%d\n", nums1_2, nums2_2, m2, n2)
	merge(nums1_2, m2, nums2_2, n2)
	expected2 := []int{1}
	fmt.Printf("After:  nums1=%v\n", nums1_2)
	fmt.Printf("Expected: %v\n", expected2)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums1_2, expected2))

	// Example 3
	fmt.Println("=== Example 3 ===")
	nums1_3 := []int{0}
	nums2_3 := []int{1}
	m3, n3 := 0, 1
	fmt.Printf("Before: nums1=%v, nums2=%v, m=%d, n=%d\n", nums1_3, nums2_3, m3, n3)
	merge(nums1_3, m3, nums2_3, n3)
	expected3 := []int{1}
	fmt.Printf("After:  nums1=%v\n", nums1_3)
	fmt.Printf("Expected: %v\n", expected3)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums1_3, expected3))

	// Test 4: All from nums2
	fmt.Println("=== Test 4: All from nums2 ===")
	nums1_4 := []int{0, 0, 0}
	nums2_4 := []int{1, 2, 3}
	m4, n4 := 0, 3
	fmt.Printf("Before: nums1=%v, nums2=%v, m=%d, n=%d\n", nums1_4, nums2_4, m4, n4)
	merge(nums1_4, m4, nums2_4, n4)
	expected4 := []int{1, 2, 3}
	fmt.Printf("After:  nums1=%v\n", nums1_4)
	fmt.Printf("Expected: %v\n", expected4)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums1_4, expected4))

	// Test 5: Interleaved
	fmt.Println("=== Test 5: Interleaved ===")
	nums1_5 := []int{1, 3, 5, 0, 0, 0}
	nums2_5 := []int{2, 4, 6}
	m5, n5 := 3, 3
	fmt.Printf("Before: nums1=%v, nums2=%v, m=%d, n=%d\n", nums1_5, nums2_5, m5, n5)
	merge(nums1_5, m5, nums2_5, n5)
	expected5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("After:  nums1=%v\n", nums1_5)
	fmt.Printf("Expected: %v\n", expected5)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums1_5, expected5))

	// Test 6: nums1 all larger
	fmt.Println("=== Test 6: nums1 all larger ===")
	nums1_6 := []int{4, 5, 6, 0, 0, 0}
	nums2_6 := []int{1, 2, 3}
	m6, n6 := 3, 3
	fmt.Printf("Before: nums1=%v, nums2=%v, m=%d, n=%d\n", nums1_6, nums2_6, m6, n6)
	merge(nums1_6, m6, nums2_6, n6)
	expected6 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("After:  nums1=%v\n", nums1_6)
	fmt.Printf("Expected: %v\n", expected6)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums1_6, expected6))

	// Test 7: Negative numbers
	fmt.Println("=== Test 7: Negative numbers ===")
	nums1_7 := []int{-3, -1, 0, 0, 0}
	nums2_7 := []int{-2, 0, 1}
	m7, n7 := 2, 3
	fmt.Printf("Before: nums1=%v, nums2=%v, m=%d, n=%d\n", nums1_7, nums2_7, m7, n7)
	merge(nums1_7, m7, nums2_7, n7)
	expected7 := []int{-3, -2, -1, 0, 1}
	fmt.Printf("After:  nums1=%v\n", nums1_7)
	fmt.Printf("Expected: %v\n", expected7)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums1_7, expected7))

	// Test 8: Duplicates
	fmt.Println("=== Test 8: Duplicates ===")
	nums1_8 := []int{1, 2, 2, 0, 0, 0}
	nums2_8 := []int{2, 2, 3}
	m8, n8 := 3, 3
	fmt.Printf("Before: nums1=%v, nums2=%v, m=%d, n=%d\n", nums1_8, nums2_8, m8, n8)
	merge(nums1_8, m8, nums2_8, n8)
	expected8 := []int{1, 2, 2, 2, 2, 3}
	fmt.Printf("After:  nums1=%v\n", nums1_8)
	fmt.Printf("Expected: %v\n", expected8)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums1_8, expected8))
}
