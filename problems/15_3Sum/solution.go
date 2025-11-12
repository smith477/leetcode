package main

import (
	"fmt"
	"sort"
)

// Two pointer pattern (but with three pointers)
// Time complexity O(n²), Space complexity O(1)
// Array needs to be sorted
// Using three pointers where one is on the beginning, one in the middle and one on the right
// The left one is fixed pointer that while the middle and right one changes based on whether the sum is lesser or larger than 0 (move middle if sum < 0, move right if sum > 0)
// When starting iteration it's important to skip all same elements
// After that assign middle to one more than left and right to one less than total length - 1
// While middle is less than right try different combinations while sum is 0
// While middle is less than right skip duplicates if sum is 0 after finding the triplet
// return result
func threeSum(nums []int) [][]int {

	result := [][]int{}

	n := len(nums)
	sort.Ints(nums)

	for left := 0; left < n-2; left++ {
		if left > 0 && nums[left-1] == nums[left] {
			continue
		}

		middle, right := left+1, n-1

		for middle < right {
			localSum := nums[left] + nums[middle] + nums[right]

			if localSum == 0 {
				result = append(result, []int{nums[left], nums[middle], nums[right]})

				for middle < right && nums[middle] == nums[middle+1] {
					middle++
				}
				for middle < right && nums[right] == nums[right-1] {
					right--
				}
				middle++
				right--
			} else if localSum < 0 {
				middle++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	// Test cases
	tests := [][]int{
		{-1, 0, 1, 2, -1, -4}, // Expect: [[-1, -1, 2], [-1, 0, 1]]
		{0, 1, 1},             // Expect: []
		{0, 0, 0},             // Expect: [[0, 0, 0]]
	}

	for i, nums := range tests {
		fmt.Printf("Test %d: Input = %v\n", i+1, nums)
		result := threeSum(nums)
		fmt.Printf("Output = %v\n\n", result)
	}
}
