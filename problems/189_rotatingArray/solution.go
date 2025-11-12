package main

import (
	"fmt"
	"reflect"
)

func rotate(nums []int, k int) {
	if k == 0 || len(nums) <= 1 {
		return
	}

	arrayCount := len(nums)
	k = k % arrayCount
	if k == 0 {
		return
	}

	// Reverse the whole array
	reverse(nums, 0, len(nums)-1)
	// Reverse first k places
	reverse(nums, 0, k-1)
	// Reverse the rest of the array
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	// Example 1
	fmt.Println("=== Example 1 ===")
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	expected1 := []int{5, 6, 7, 1, 2, 3, 4}
	fmt.Printf("Before: %v, k=%d\n", nums1, k1)
	rotate(nums1, k1)
	fmt.Printf("After:  %v\n", nums1)
	fmt.Printf("Expected: %v\n", expected1)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums1, expected1))

	// Example 2
	fmt.Println("=== Example 2 ===")
	nums2 := []int{-1, -100, 3, 99}
	k2 := 2
	expected2 := []int{3, 99, -1, -100}
	fmt.Printf("Before: %v, k=%d\n", nums2, k2)
	rotate(nums2, k2)
	fmt.Printf("After:  %v\n", nums2)
	fmt.Printf("Expected: %v\n", expected2)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums2, expected2))

	// Test 3: k larger than array length
	fmt.Println("=== Test 3: k > n ===")
	nums3 := []int{1, 2, 3, 4, 5}
	k3 := 7 // Same as k=2
	expected3 := []int{4, 5, 1, 2, 3}
	fmt.Printf("Before: %v, k=%d\n", nums3, k3)
	rotate(nums3, k3)
	fmt.Printf("After:  %v\n", nums3)
	fmt.Printf("Expected: %v\n", expected3)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums3, expected3))

	// Test 4: k = 0
	fmt.Println("=== Test 4: k = 0 ===")
	nums4 := []int{1, 2, 3}
	k4 := 0
	expected4 := []int{1, 2, 3}
	fmt.Printf("Before: %v, k=%d\n", nums4, k4)
	rotate(nums4, k4)
	fmt.Printf("After:  %v\n", nums4)
	fmt.Printf("Expected: %v\n", expected4)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums4, expected4))

	// Test 5: Single element
	fmt.Println("=== Test 5: Single Element ===")
	nums5 := []int{1}
	k5 := 100
	expected5 := []int{1}
	fmt.Printf("Before: %v, k=%d\n", nums5, k5)
	rotate(nums5, k5)
	fmt.Printf("After:  %v\n", nums5)
	fmt.Printf("Expected: %v\n", expected5)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums5, expected5))

	// Test 6: k = n (full rotation)
	fmt.Println("=== Test 6: k = n ===")
	nums6 := []int{1, 2, 3, 4}
	k6 := 4
	expected6 := []int{1, 2, 3, 4}
	fmt.Printf("Before: %v, k=%d\n", nums6, k6)
	rotate(nums6, k6)
	fmt.Printf("After:  %v\n", nums6)
	fmt.Printf("Expected: %v\n", expected6)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums6, expected6))

	// Test 7: Two elements
	fmt.Println("=== Test 7: Two Elements ===")
	nums7 := []int{1, 2}
	k7 := 1
	expected7 := []int{2, 1}
	fmt.Printf("Before: %v, k=%d\n", nums7, k7)
	rotate(nums7, k7)
	fmt.Printf("After:  %v\n", nums7)
	fmt.Printf("Expected: %v\n", expected7)
	fmt.Printf("Match: %v\n\n", reflect.DeepEqual(nums7, expected7))
}
