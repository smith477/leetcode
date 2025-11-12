package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {

	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for {
		slow = nums[slow]
		fast = nums[fast]

		if slow == fast {
			break
		}
	}

	return slow
}

func main() {
	// Test cases
	tests := [][]int{
		{1, 3, 4, 2, 2},
		{3, 1, 3, 4, 2},
		{3, 3, 3, 3, 3},
	}

	for _, nums := range tests {
		fmt.Printf("Input: %v -> Duplicate: %d\n", nums, findDuplicate(nums))
	}
}
