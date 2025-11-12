package main

import "fmt"

// Binary Search on Rotated Sorted Array
// Time complexity O(log(n)), Space complexity O(1)
// Find middle element: start + (end-start)/2 to avoid integer overflow
// If middle is target, return it
// Determine which half is properly sorted (left or right)
// If left half is sorted (nums[start] <= nums[mid]):
//   - Check if target is within sorted left half range
//   - If yes, search left (end = mid-1), else search right (start = mid+1)
//
// If right half is sorted:
//   - Check if target is within sorted right half range
//   - If yes, search right (start = mid+1), else search left (end = mid-1)
//
// Repeat until target found or start > end
func search(nums []int, target int) int {

	start, end := 0, len(nums)-1

	var mid = 0
	for start <= end {

		mid = start + (end-start)/2

		if nums[mid] == target {
			return mid
		}

		if nums[start] <= nums[mid] {
			if nums[start] <= target && nums[mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && nums[end] >= target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

func main() {
	// Test cases from the problem

	// Example 1
	nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	target1 := 0
	result1 := search(nums1, target1)
	fmt.Printf("Example 1: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("Output: %d (Expected: 4)\n\n", result1)

	// Example 2
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 3
	result2 := search(nums2, target2)
	fmt.Printf("Example 2: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("Output: %d (Expected: -1)\n\n", result2)

	// Example 3
	nums3 := []int{1}
	target3 := 0
	result3 := search(nums3, target3)
	fmt.Printf("Example 3: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("Output: %d (Expected: -1)\n\n", result3)

	// Additional test cases
	nums4 := []int{4, 5, 6, 7, 0, 1, 2}
	target4 := 5
	result4 := search(nums4, target4)
	fmt.Printf("Additional Test: nums = %v, target = %d\n", nums4, target4)
	fmt.Printf("Output: %d (Expected: 1)\n\n", result4)

	// No rotation
	nums5 := []int{1, 2, 3, 4, 5}
	target5 := 3
	result5 := search(nums5, target5)
	fmt.Printf("No rotation: nums = %v, target = %d\n", nums5, target5)
	fmt.Printf("Output: %d (Expected: 2)\n\n", result5)
}
