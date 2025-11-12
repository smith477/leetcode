package main

import "fmt"

func sortColors(nums []int) []int {

	n := len(nums)
	if n <= 1 {
		return nums
	}

	left, mid, right := 0, 0, n-1

	for mid < right {
		switch nums[mid] {
		case 0:
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[right] = nums[right], nums[mid]
			right--
		}
	}

	return nums
}

func main() {
	tests := [][]int{
		{2, 0, 0, 1, 2, 1},
		{2, 2, 2, 1, 1, 0, 0},
		{0, 1, 2},
		{1, 0, 2, 1, 0, 2},
		{0, 0, 0},
		{2, 1, 0},
		{},
		{1},
		{0, 2, 2, 1, 0, 1},
	}

	for _, test := range tests {
		fmt.Printf("Before: %v\n", test)
		sortColors(test)
		fmt.Printf("After:  %v\n\n", test)
	}
}
