package main

import "fmt"

// Binary Search - Find First and Last Position
// Time complexity O(log(n)), Space complexity O(1)
// Use helper method findBound() twice: once for left bound (isFirst=true), once for right bound (isFirst=false)
// In findBound(): use standard binary search, but when target is found:
//   - For left bound: continue searching left half (end = mid - 1) to find first occurrence
//   - For right bound: continue searching right half (start = mid + 1) to find last occurrence
//
// If first occurrence not found, return [-1, -1]
// Otherwise return [firstOccurrence, lastOccurrence]
func searchRange(nums []int, target int) []int {
	if len(nums) < 2 {
	}

	leftOccurence := findOccurence(nums, target, true)

	if leftOccurence == -1 {
		return []int{-1, -1}
	}

	rightOccurence := findOccurence(nums, target, false)

	return []int{leftOccurence, rightOccurence}
}

func findOccurence(nums []int, target int, isFirst bool) int {
	start, end := 0, len(nums)-1
	result := -1

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			result = mid
			if isFirst {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return result
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1, -1]
	fmt.Println(searchRange([]int{}, 0))                  // [-1, -1]
}
