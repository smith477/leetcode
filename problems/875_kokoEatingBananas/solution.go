package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {

	left, right := 1, 0

	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}

	var mid = 0

	for left < right {
		mid = left + (right-left)/2

		if canEatPile(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canEatPile(piles []int, h int, k int) bool {

	hours := 0
	for _, pile := range piles {
		hours += (pile + k - 1) / k
		if hours > h {
			return false
		}
	}
	return true
}

func main() {
	// Example 1
	piles1 := []int{3, 6, 7, 11}
	h1 := 8
	fmt.Printf("Example 1:\n")
	fmt.Printf("Input: piles = %v, h = %d\n", piles1, h1)
	fmt.Printf("Output: %d\n", minEatingSpeed(piles1, h1))
	fmt.Printf("Expected: 4\n\n")

	// Example 2
	piles2 := []int{30, 11, 23, 4, 20}
	h2 := 5
	fmt.Printf("Example 2:\n")
	fmt.Printf("Input: piles = %v, h = %d\n", piles2, h2)
	fmt.Printf("Output: %d\n", minEatingSpeed(piles2, h2))
	fmt.Printf("Expected: 30\n\n")

	// Example 3
	piles3 := []int{30, 11, 23, 4, 20}
	h3 := 6
	fmt.Printf("Example 3:\n")
	fmt.Printf("Input: piles = %v, h = %d\n", piles3, h3)
	fmt.Printf("Output: %d\n", minEatingSpeed(piles3, h3))
	fmt.Printf("Expected: 23\n\n")

	// Additional test cases
	piles4 := []int{1000000000}
	h4 := 2
	fmt.Printf("Edge Case - Large pile:\n")
	fmt.Printf("Input: piles = %v, h = %d\n", piles4, h4)
	fmt.Printf("Output: %d\n", minEatingSpeed(piles4, h4))
	fmt.Printf("Expected: 500000000\n\n")

	piles5 := []int{1, 1, 1, 1}
	h5 := 4
	fmt.Printf("Edge Case - All ones:\n")
	fmt.Printf("Input: piles = %v, h = %d\n", piles5, h5)
	fmt.Printf("Output: %d\n", minEatingSpeed(piles5, h5))
	fmt.Printf("Expected: 1\n\n")
}
