package main

import "fmt"

// Two-pointer approach allowing up to 2 duplicates
// Time: O(n), Space: O(1)
// Use k to track position for next valid element
// Compare current element with element at k-2, if different then we can include it
// This ensures at most 2 consecutive duplicates (current can match k-1 but not k-2)
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	k := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	expectedNums := []int{1, 1, 2, 2, 3}

	k := removeDuplicates(nums)

	if k != len(expectedNums) {
		fmt.Printf("Test failed! Expected length %d but got %d\n", len(expectedNums), k)
		return
	}

	for i := 0; i < k; i++ {
		if nums[i] != expectedNums[i] {
			fmt.Printf("Test failed! At index %d, expected %d but got %d\n", i, expectedNums[i], nums[i])
			return
		}
	}

	fmt.Println("Test passed!")
}
