package main

import "fmt"

func orderAgnosticSearch(arr []int, key int) int {

	if len(arr) < 1 {
		return -1
	}

	start, end := 0, len(arr)-1

	isAscending := arr[start] <= arr[end]

	var mid = 0
	for start <= end {
		mid = start + (end-start)/2

		if arr[mid] == key {
			return mid
		}

		if isAscending {
			if arr[mid] > key {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if arr[mid] < key {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}

type inputType struct {
	nums []int
	key  int
}

type TestCase struct {
	input  inputType
	output int
}

func main() {
	testCases := []TestCase{
		// Example 1: Ascending order
		{
			input: inputType{
				nums: []int{4, 6, 10},
				key:  10,
			},
			output: 2,
		},
		// Example 2: Ascending order
		{
			input: inputType{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				key:  5,
			},
			output: 4,
		},
		// Example 3: Descending order
		{
			input: inputType{
				nums: []int{10, 6, 4},
				key:  10,
			},
			output: 0,
		},
		// Example 4: Descending order
		{
			input: inputType{
				nums: []int{10, 6, 4},
				key:  4,
			},
			output: 2,
		},
		// Additional test cases
		// Single element - found
		{
			input: inputType{
				nums: []int{5},
				key:  5,
			},
			output: 0,
		},
		// Single element - not found
		{
			input: inputType{
				nums: []int{5},
				key:  3,
			},
			output: -1,
		},
		// Two elements ascending
		{
			input: inputType{
				nums: []int{1, 5},
				key:  5,
			},
			output: 1,
		},
		// Two elements descending
		{
			input: inputType{
				nums: []int{5, 1},
				key:  5,
			},
			output: 0,
		},
		// Not found in ascending
		{
			input: inputType{
				nums: []int{1, 3, 5, 7, 9},
				key:  6,
			},
			output: -1,
		},
		// Not found in descending
		{
			input: inputType{
				nums: []int{9, 7, 5, 3, 1},
				key:  6,
			},
			output: -1,
		},
		// With duplicates ascending
		{
			input: inputType{
				nums: []int{1, 2, 2, 2, 3, 4, 5},
				key:  2,
			},
			output: 1, // or 2 or 3 - any valid index of 2
		},
		// With duplicates descending
		{
			input: inputType{
				nums: []int{5, 4, 3, 3, 3, 2, 1},
				key:  3,
			},
			output: 2, // or 3 or 4 - any valid index of 3
		},
	}

	for i, tc := range testCases {
		result := orderAgnosticSearch(tc.input.nums, tc.input.key)

		// For duplicates, check if the value at result index equals key
		passed := false
		if result == -1 && tc.output == -1 {
			passed = true
		} else if result != -1 && result < len(tc.input.nums) && tc.input.nums[result] == tc.input.key {
			passed = true
		}

		if passed {
			fmt.Printf("Test %2d: PASS - Expected index %d, Got %d\n", i+1, tc.output, result)
		} else {
			fmt.Printf("Test %2d: FAIL - Expected index %d, Got %d\n", i+1, tc.output, result)
		}
	}
}
