package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {

	start, end := 0, len(arr)-1

	for start < end {
		mid := start + (end-start)/2

		if arr[mid] > arr[mid+1] {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func main() {
	// Test cases from the problem
	testCases := [][]int{
		{0, 1, 0},                     // Example 1
		{0, 2, 1, 0},                  // Example 2
		{0, 10, 5, 2},                 // Example 3
		{1, 3, 5, 7, 6, 4, 2},         // Additional example
		{5, 10, 15, 20, 18, 12, 8, 3}, // Additional example
		{3, 5, 3},                     // Additional example
		{1, 2, 3, 4, 5, 4, 3, 2, 1},   // Additional example
		{0, 5, 3, 2, 1},               // Additional example
	}

	for i, arr := range testCases {
		peakIndex := peakIndexInMountainArray(arr)
		peakValue := arr[peakIndex]
		fmt.Printf("Test %d: arr = %v\n", i+1, arr)
		fmt.Printf("  Peak index: %d, Peak value: %d\n\n", peakIndex, peakValue)
	}
}
