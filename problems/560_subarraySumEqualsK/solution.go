package main

import "fmt"

func subarraySum(nums []int, k int) int {

	count := 0
	currentSum := 0

	numMap := map[int]int{0: 1}

	for _, number := range nums {

		currentSum += number

		if value, found := numMap[currentSum-k]; found {
			count += value
		}

		numMap[currentSum]++
	}

	return count
}

func main() {
	// Example 1
	nums1 := []int{1, 1, 1}
	k1 := 2
	fmt.Printf("Input: nums = %v, k = %d\n", nums1, k1)
	fmt.Printf("Output: %d\n\n", subarraySum(nums1, k1))

	// Example 2
	nums2 := []int{1, 2, 3}
	k2 := 3
	fmt.Printf("Input: nums = %v, k = %d\n", nums2, k2)
	fmt.Printf("Output: %d\n", subarraySum(nums2, k2))
}
