package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {

	prod := 1
	left := -1
	result := 0

	for i := 0; i < len(nums); i++ {

		prod *= nums[i]

		for prod >= k {
			left++
			prod /= nums[left]
		}
		result += i - left

	}

	return result

}

func main() {
	tests := []struct {
		nums   []int
		k      int
		output int
	}{
		{[]int{10, 5, 2, 6}, 100, 8},
		{[]int{1, 2, 3}, 0, 0},
		{[]int{1, 1, 1}, 2, 6},
		{[]int{1, 2, 3, 4}, 10, 7},
	}

	for _, test := range tests {
		result := numSubarrayProductLessThanK(test.nums, test.k)
		fmt.Printf("Input: nums=%v, k=%d -> Output: %d (Expected: %d)\n", test.nums, test.k, result, test.output)
	}
}
