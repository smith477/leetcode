package main

import (
	"fmt"
	"math"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Two pointer pattern (but with three pointers)
// Time complexity O(n²), Space complexity O(1)
// First check if there are enough elements
// Array must be sorted so you can have left and right of the array
// Start iterating
// From left pointer all repeating elements should be skipped
// Initialize middle and right pointer with left + 1 and last element of the array
// Start moving middle pointer towards right one to calculate all combinations
// Calculate local sum to find closer sum to target
// Based on local sum/target ration move the middle or right pointer
// If sum is lower than target that means middle pointer should be increased otherwise right pointer should be decreased
// If sum is same as target then that sum is the result
func threeSumClosest(nums []int, target int) int {

	if len(nums) <= 2 {
		return 0
	}

	n := len(nums)

	sort.Ints(nums)

	globalSum := nums[0] + nums[1] + nums[2]

	minDif := math.MaxInt32
	for left := 0; left < n-2; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}

		middle, right := left+1, n-1

		for middle < right {
			localSum := nums[left] + nums[middle] + nums[right]

			absSum := abs(localSum - target)
			if minDif > absSum {
				minDif = absSum
				globalSum = localSum
			}

			if localSum == target {
				return localSum
			} else if localSum > target {
				right--
			} else {
				middle++
			}
		}
	}

	return globalSum
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	result := threeSumClosest(nums, target)
	fmt.Println("Closest sum:", result)
}
