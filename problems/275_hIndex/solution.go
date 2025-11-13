package main

import (
	"fmt"
	"sort"
)

// func hIndex(citations []int) int {

// 	maxH := len(citations)
// 	sort.Ints(citations)

// 	for _, citation := range citations {

// 		if citation < maxH {
// 			maxH--
// 		}
// 	}

// 	return maxH
// }

func hIndex(citations []int) int {

	h := 0
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for i := 0; i < len(citations); i++ {
		if citations[i] >= i+1 {
			h++
		}
	}

	return h
}

func main() {
	// Example 1
	fmt.Println("=== Example 1 ===")
	citations1 := []int{3, 0, 6, 1, 5}
	result1 := hIndex(citations1)
	fmt.Printf("Input: %v\n", citations1)
	fmt.Printf("Output: %d\n", result1)
	fmt.Printf("Expected: 3\n")
	fmt.Printf("Explanation: 3 papers with at least 3 citations each\n")
	fmt.Printf("Match: %v\n\n", result1 == 3)

	// Example 2
	fmt.Println("=== Example 2 ===")
	citations2 := []int{1, 3, 1}
	result2 := hIndex(citations2)
	fmt.Printf("Input: %v\n", citations2)
	fmt.Printf("Output: %d\n", result2)
	fmt.Printf("Expected: 1\n")
	fmt.Printf("Explanation: 1 paper with at least 1 citation\n")
	fmt.Printf("Match: %v\n\n", result2 == 1)

	// Test 3: All zeros
	fmt.Println("=== Test 3: All Zeros ===")
	citations3 := []int{0, 0, 0, 0}
	result3 := hIndex(citations3)
	fmt.Printf("Input: %v\n", citations3)
	fmt.Printf("Output: %d\n", result3)
	fmt.Printf("Expected: 0\n")
	fmt.Printf("Match: %v\n\n", result3 == 0)

	// Test 4: Single paper with many citations
	fmt.Println("=== Test 4: Single Paper ===")
	citations4 := []int{100}
	result4 := hIndex(citations4)
	fmt.Printf("Input: %v\n", citations4)
	fmt.Printf("Output: %d\n", result4)
	fmt.Printf("Expected: 1\n")
	fmt.Printf("Explanation: Max h-index is 1 (only 1 paper)\n")
	fmt.Printf("Match: %v\n\n", result4 == 1)

	// Test 5: All high citations
	fmt.Println("=== Test 5: All High Citations ===")
	citations5 := []int{10, 8, 5, 4, 3}
	result5 := hIndex(citations5)
	fmt.Printf("Input: %v\n", citations5)
	fmt.Printf("Output: %d\n", result5)
	fmt.Printf("Expected: 4\n")
	fmt.Printf("Explanation: 4 papers with at least 4 citations\n")
	fmt.Printf("Match: %v\n\n", result5 == 4)

	// Test 6: Increasing order
	fmt.Println("=== Test 6: Increasing Order ===")
	citations6 := []int{1, 2, 3, 4, 5}
	result6 := hIndex(citations6)
	fmt.Printf("Input: %v\n", citations6)
	fmt.Printf("Output: %d\n", result6)
	fmt.Printf("Expected: 3\n")
	fmt.Printf("Explanation: 3 papers with at least 3 citations\n")
	fmt.Printf("Match: %v\n\n", result6 == 3)

	// Test 7: All same citations
	fmt.Println("=== Test 7: All Same ===")
	citations7 := []int{5, 5, 5, 5, 5}
	result7 := hIndex(citations7)
	fmt.Printf("Input: %v\n", citations7)
	fmt.Printf("Output: %d\n", result7)
	fmt.Printf("Expected: 5\n")
	fmt.Printf("Explanation: 5 papers with at least 5 citations\n")
	fmt.Printf("Match: %v\n\n", result7 == 5)

	// Test 8: Mix of high and low
	fmt.Println("=== Test 8: Mix ===")
	citations8 := []int{100, 2, 1}
	result8 := hIndex(citations8)
	fmt.Printf("Input: %v\n", citations8)
	fmt.Printf("Output: %d\n", result8)
	fmt.Printf("Expected: 2\n")
	fmt.Printf("Explanation: 2 papers with at least 2 citations\n")
	fmt.Printf("Match: %v\n\n", result8 == 2)

	// Test 9: Large h-index
	fmt.Println("=== Test 9: Large h-index ===")
	citations9 := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	result9 := hIndex(citations9)
	fmt.Printf("Input: %v\n", citations9)
	fmt.Printf("Output: %d\n", result9)
	fmt.Printf("Expected: 10\n")
	fmt.Printf("Match: %v\n\n", result9 == 10)

	// Test 10: Edge case
	fmt.Println("=== Test 10: Edge Case ===")
	citations10 := []int{0, 1, 3, 5, 6}
	result10 := hIndex(citations10)
	fmt.Printf("Input: %v\n", citations10)
	fmt.Printf("Output: %d\n", result10)
	fmt.Printf("Expected: 3\n")
	fmt.Printf("Explanation: Papers sorted [0,1,3,5,6] - 3 papers with ≥3 citations\n")
	fmt.Printf("Match: %v\n\n", result10 == 3)

	// Test 11: Two papers
	fmt.Println("=== Test 11: Two Papers ===")
	citations11 := []int{4, 4}
	result11 := hIndex(citations11)
	fmt.Printf("Input: %v\n", citations11)
	fmt.Printf("Output: %d\n", result11)
	fmt.Printf("Expected: 2\n")
	fmt.Printf("Match: %v\n\n", result11 == 2)

	// Test 12: Complex case
	fmt.Println("=== Test 12: Complex Case ===")
	citations12 := []int{11, 15, 7, 8, 12}
	result12 := hIndex(citations12)
	fmt.Printf("Input: %v\n", citations12)
	fmt.Printf("Output: %d\n", result12)
	fmt.Printf("Expected: 5\n")
	fmt.Printf("Explanation: All 5 papers have at least 5 citations\n")
	fmt.Printf("Match: %v\n\n", result12 == 5)
}
