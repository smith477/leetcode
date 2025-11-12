package main

import (
	"fmt"
	"sort"
)

func threeSumSmaller(nums []int, target int) int {

	n := len(nums)
	if n < 2 {
		return 0
	}

	result := 0

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if target > sum {
				result += right - left
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{[]int{-2, 0, 1}, 2, 1},      // [-2, 0, 1]
		{[]int{-2, 0, -1, 3}, 2, 3},  // [-2, 0, -1], [-2, 0, 3], [-2, -1, 3]
		{[]int{-1, 0, 1, 2}, 2, 2},   // [-1, 0, 1], [-1, 0, 2]
		{[]int{3, 1, 0, -2}, 4, 3},   // [-2, 0, 1], [-2, 0, 3], [-2, 1, 3]
		{[]int{0, 0, 0}, 1, 1},       // [0, 0, 0]
		{[]int{1, 1, -2}, 2, 1},      // [-2, 1, 1]
		{[]int{5, 5, 5, 5}, 10, 0},   // none
		{[]int{-5, -2, -1}, 0, 1},    // [-5, -2, -1]
		{[]int{1, 2, 3, 4, 5}, 9, 4}, // [1,2,3], [1,2,4], [1,3,4], [2,3,4]
	}

	for i, test := range tests {
		result := threeSumSmaller(test.nums, test.target)
		status := "❌"
		if result == test.output {
			status = "✅"
		}
		fmt.Printf("Test %d: Got %d | Expected %d [%s]\n", i+1, result, test.output, status)
	}
}
