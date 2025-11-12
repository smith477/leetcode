package main

import "fmt"

func findCeiling(arr []int, key int) int {
	if key > arr[len(arr)-1] {
		return -1
	}

	start, end := 0, len(arr)-1

	var mid = 0
	for start <= end {
		mid = start + (end-start)/2

		if arr[mid] < key {
			start = mid + 1
		} else if arr[mid] > key {
			end = mid - 1
		} else {
			return mid
		}
	}

	return start
}

func main() {
	// Test cases from the problem
	testCases := []struct {
		arr      []int
		key      int
		expected int
	}{
		{[]int{4, 6, 10}, 6, 1},
		{[]int{1, 3, 8, 10, 15}, 12, 4},
		{[]int{4, 6, 10}, 17, -1},
		{[]int{4, 6, 10}, -1, 0},
	}

	fmt.Println("Testing Ceiling of a Number:\n")

	for i, tc := range testCases {
		result := findCeiling(tc.arr, tc.key)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}

		fmt.Printf("Example %d:\n", i+1)
		fmt.Printf("  Input: %v, key = %d\n", tc.arr, tc.key)
		fmt.Printf("  Output: %d\n", result)
		fmt.Printf("  Expected: %d %s\n", tc.expected, status)

		if result != -1 {
			fmt.Printf("  Explanation: The smallest number >= '%d' is '%d' at index '%d'\n",
				tc.key, tc.arr[result], result)
		} else {
			fmt.Printf("  Explanation: No number >= '%d' exists in the array\n", tc.key)
		}
		fmt.Println()
	}
}
