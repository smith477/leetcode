package main

import "fmt"

// Two pointer
// Time complexity O(n), Space complexity O(1)
// Pointer k tracks where to write next distinct (non-val) element
// Iterate through array with i
// When nums[i] != val, write it at position k and increment k
// When nums[i] == val, skip it (only i increments, k stays same)
// This way k only advances when writing non-val elements
// Return k as the count of non-val elements
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	// Example 1
	fmt.Println("=== Example 1 ===")
	nums1 := []int{3, 2, 2, 3}
	val1 := 3
	k1 := removeElement(nums1, val1)
	fmt.Printf("k = %d, nums = %v (first %d elements)\n", k1, nums1[:k1], k1)
	fmt.Println("Expected: k = 2, nums = [2,2]")

	// Example 2
	fmt.Println("\n=== Example 2 ===")
	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	k2 := removeElement(nums2, val2)
	fmt.Printf("k = %d, nums = %v (first %d elements)\n", k2, nums2[:k2], k2)
	fmt.Println("Expected: k = 5, nums = [0,1,3,0,4]")

	// Test 3: All same value
	fmt.Println("\n=== Test 3: All Same ===")
	nums3 := []int{2, 2, 2, 2}
	val3 := 2
	k3 := removeElement(nums3, val3)
	fmt.Printf("k = %d, nums = %v\n", k3, nums3)
	fmt.Println("Expected: k = 0")

	// Test 4: No matches
	fmt.Println("\n=== Test 4: No Matches ===")
	nums4 := []int{1, 2, 3, 4}
	val4 := 5
	k4 := removeElement(nums4, val4)
	fmt.Printf("k = %d, nums = %v\n", k4, nums4[:k4])
	fmt.Println("Expected: k = 4, nums = [1,2,3,4]")

	// Test 5: Empty array
	fmt.Println("\n=== Test 5: Empty ===")
	nums5 := []int{}
	val5 := 1
	k5 := removeElement(nums5, val5)
	fmt.Printf("k = %d\n", k5)
	fmt.Println("Expected: k = 0")
}
