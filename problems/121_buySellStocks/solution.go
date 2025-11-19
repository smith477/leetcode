package main

import "fmt"

// Greedy single-pass approach
// Time: O(n), Space: O(1)
// Track minimum price seen so far and maximum profit
// For each price, calculate profit if sold today and update max profit
// Update minimum price for future calculations
func maxProfit(prices []int) int {

	maxProfit := 0
	minPrice := prices[0]

	for _, price := range prices[1:] {
		profit := price - minPrice
		maxProfit = max(maxProfit, profit)
		minPrice = min(minPrice, price)
	}

	return maxProfit
}

func main() {
	// Test cases
	testCases := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{3, 3, 3, 3}, 0},
		{[]int{2, 4, 1}, 2},
		{[]int{1}, 0},
		{[]int{}, 0},
	}

	fmt.Println("Running test cases...\n")

	for i, tc := range testCases {
		result := maxProfit(tc.prices)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("Test %d: %s\n", i+1, status)
		fmt.Printf("  Input:    %v\n", tc.prices)
		fmt.Printf("  Expected: %d\n", tc.expected)
		fmt.Printf("  Got:      %d\n", result)
		fmt.Println()
	}
}
