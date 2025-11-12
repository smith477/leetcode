package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {

	start, end := 0, len(letters)-1

	var mid = 0
	for start <= end {
		mid = start + (end-start)/2

		if letters[mid] <= target {
			start = mid + 1
		} else {
			end = mid
		}
	}

	if letters[start] > target {
		return letters[start]
	}

	return letters[0]
}

func main() {
	// Test Example 1
	letters1 := []byte{'c', 'f', 'j'}
	target1 := byte('a')
	result1 := nextGreatestLetter(letters1, target1)
	fmt.Printf("Example 1: Input: letters = %c, target = '%c'\n", letters1, target1)
	fmt.Printf("Output: '%c'\n\n", result1)

	// Test Example 2
	letters2 := []byte{'c', 'f', 'j'}
	target2 := byte('c')
	result2 := nextGreatestLetter(letters2, target2)
	fmt.Printf("Example 2: Input: letters = %c, target = '%c'\n", letters2, target2)
	fmt.Printf("Output: '%c'\n\n", result2)

	// Test Example 3
	letters3 := []byte{'x', 'x', 'y', 'y'}
	target3 := byte('z')
	result3 := nextGreatestLetter(letters3, target3)
	fmt.Printf("Example 3: Input: letters = %c, target = '%c'\n", letters3, target3)
	fmt.Printf("Output: '%c'\n\n", result3)
}
