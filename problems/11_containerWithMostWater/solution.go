package main

import (
	"fmt"
	"math"
)

// Two-pointer greedy approach
// Time: O(n), Space: O(1)
// Start with widest container, move pointer with shorter height inward to find taller lines
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := math.MinInt

	for left < right {
		width := right - left
		currentArea := width * min(height[left], height[right])
		maxWater = max(currentArea, maxWater)

		// Move pointer with shorter height
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func main() {
	// Test case 1
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("Test 1: %v\n", height1)
	fmt.Printf("Output: %d\n", maxArea(height1))
	fmt.Printf("Expected: 49\n\n")

	// Test case 2
	height2 := []int{1, 1}
	fmt.Printf("Test 2: %v\n", height2)
	fmt.Printf("Output: %d\n", maxArea(height2))
	fmt.Printf("Expected: 1\n\n")

	// Test case 3
	height3 := []int{4, 3, 2, 1, 4}
	fmt.Printf("Test 3: %v\n", height3)
	fmt.Printf("Output: %d\n", maxArea(height3))
	fmt.Printf("Expected: 16\n\n")

	// Test case 4
	height4 := []int{1, 2, 1}
	fmt.Printf("Test 4: %v\n", height4)
	fmt.Printf("Output: %d\n", maxArea(height4))
	fmt.Printf("Expected: 2\n\n")
}
