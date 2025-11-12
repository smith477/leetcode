package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1
	pos := n - 1

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]

		if leftSq > rightSq {
			res[pos] = leftSq
			left++
		} else {
			res[pos] = rightSq
			right--
		}
		pos--
	}
	return res
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	result := sortedSquares(nums)
	fmt.Println(result) // Output: [0 1 9 16 100]
}
