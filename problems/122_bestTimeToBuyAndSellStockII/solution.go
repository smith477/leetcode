package main

import "fmt"

// Greedy approach - capture all upward price movements
// Time: O(n), Space: O(1)
// Sum all positive price differences (buy before each increase, sell at peak)
// Equivalent to buying at every valley and selling at every peak
func maxProfit(prices []int) int {

	totalProfit := 0

	for i := 1; i < len(prices); i++ {
		totalProfit += max(0, prices[i]-prices[i-1])
	}

	return totalProfit
}

func main() {
	// Example 1
	fmt.Println("=== Example 1 ===")
	prices1 := []int{7, 1, 5, 3, 6, 4}
	result1 := maxProfit(prices1)
	fmt.Printf("Input: %v\n", prices1)
	fmt.Printf("Output: %d\n", result1)
	fmt.Printf("Expected: 7\n")
	fmt.Printf("Explanation: Buy at 1, sell at 5 (profit=4), then buy at 3, sell at 6 (profit=3). Total=7\n")
	fmt.Printf("Match: %v\n\n", result1 == 7)

	// Example 2
	fmt.Println("=== Example 2 ===")
	prices2 := []int{1, 2, 3, 4, 5}
	result2 := maxProfit(prices2)
	fmt.Printf("Input: %v\n", prices2)
	fmt.Printf("Output: %d\n", result2)
	fmt.Printf("Expected: 4\n")
	fmt.Printf("Explanation: Buy at 1, sell at 5. Profit = 5-1 = 4\n")
	fmt.Printf("Match: %v\n\n", result2 == 4)

	// Example 3
	fmt.Println("=== Example 3 ===")
	prices3 := []int{7, 6, 4, 3, 1}
	result3 := maxProfit(prices3)
	fmt.Printf("Input: %v\n", prices3)
	fmt.Printf("Output: %d\n", result3)
	fmt.Printf("Expected: 0\n")
	fmt.Printf("Explanation: Prices only decrease, no profit possible\n")
	fmt.Printf("Match: %v\n\n", result3 == 0)

	// Test 4: Single element
	fmt.Println("=== Test 4: Single Element ===")
	prices4 := []int{5}
	result4 := maxProfit(prices4)
	fmt.Printf("Input: %v\n", prices4)
	fmt.Printf("Output: %d\n", result4)
	fmt.Printf("Expected: 0\n")
	fmt.Printf("Match: %v\n\n", result4 == 0)

	// Test 5: Two elements - profit
	fmt.Println("=== Test 5: Two Elements (Profit) ===")
	prices5 := []int{1, 5}
	result5 := maxProfit(prices5)
	fmt.Printf("Input: %v\n", prices5)
	fmt.Printf("Output: %d\n", result5)
	fmt.Printf("Expected: 4\n")
	fmt.Printf("Match: %v\n\n", result5 == 4)

	// Test 6: Two elements - no profit
	fmt.Println("=== Test 6: Two Elements (No Profit) ===")
	prices6 := []int{5, 1}
	result6 := maxProfit(prices6)
	fmt.Printf("Input: %v\n", prices6)
	fmt.Printf("Output: %d\n", result6)
	fmt.Printf("Expected: 0\n")
	fmt.Printf("Match: %v\n\n", result6 == 0)

	// Test 7: Multiple peaks and valleys
	fmt.Println("=== Test 7: Multiple Peaks and Valleys ===")
	prices7 := []int{3, 1, 4, 1, 5, 9, 2, 6}
	result7 := maxProfit(prices7)
	fmt.Printf("Input: %v\n", prices7)
	fmt.Printf("Output: %d\n", result7)
	// (4-1) + (5-1) + (9-5) + (6-2) = 3 + 4 + 4 + 4 = 15
	fmt.Printf("Expected: 15\n")
	fmt.Printf("Explanation: Capture all upward moves: (4-1)+(5-1)+(9-5)+(6-2) = 3+4+4+4 = 15\n")
	fmt.Printf("Match: %v\n\n", result7 == 15)

	// Test 8: All same prices
	fmt.Println("=== Test 8: All Same Prices ===")
	prices8 := []int{5, 5, 5, 5, 5}
	result8 := maxProfit(prices8)
	fmt.Printf("Input: %v\n", prices8)
	fmt.Printf("Output: %d\n", result8)
	fmt.Printf("Expected: 0\n")
	fmt.Printf("Match: %v\n\n", result8 == 0)

	// Test 9: Zigzag pattern
	fmt.Println("=== Test 9: Zigzag Pattern ===")
	prices9 := []int{1, 3, 2, 4, 3, 5}
	result9 := maxProfit(prices9)
	fmt.Printf("Input: %v\n", prices9)
	fmt.Printf("Output: %d\n", result9)
	// (3-1) + (4-2) + (5-3) = 2 + 2 + 2 = 6
	fmt.Printf("Expected: 6\n")
	fmt.Printf("Match: %v\n\n", result9 == 6)

	// Test 10: Large numbers
	fmt.Println("=== Test 10: Large Numbers ===")
	prices10 := []int{100, 180, 260, 310, 40, 535, 695}
	result10 := maxProfit(prices10)
	fmt.Printf("Input: %v\n", prices10)
	fmt.Printf("Output: %d\n", result10)
	// (180-100) + (260-180) + (310-260) + (535-40) + (695-535)
	// = 80 + 80 + 50 + 495 + 160 = 865
	fmt.Printf("Expected: 865\n")
	fmt.Printf("Match: %v\n\n", result10 == 865)
}
