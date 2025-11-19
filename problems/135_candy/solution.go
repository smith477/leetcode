package main

import "fmt"

func candy(ratings []int) int {

	n := len(ratings)
	result := make([]int, len(ratings))

	for i := 0; i < len(ratings); i++ {
		result[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			result[i] = result[i-1] + 1
		}
	}

	// Right to left pass: handle decreasing ratings
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			result[i] = max(result[i], result[i+1]+1)
		}
	}

	count := 0
	for i := 0; i < len(ratings); i++ {
		count += result[i]
	}

	return count
}

func main() {
	testCases := []struct {
		name     string
		ratings  []int
		expected int
	}{
		{
			name:     "Example 1",
			ratings:  []int{1, 0, 2},
			expected: 5,
		},
		{
			name:     "Example 2",
			ratings:  []int{1, 2, 2},
			expected: 4,
		},
		{
			name:     "Single child",
			ratings:  []int{1},
			expected: 1,
		},
		{
			name:     "All same ratings",
			ratings:  []int{1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "Increasing ratings",
			ratings:  []int{1, 2, 3, 4, 5},
			expected: 15, // 1+2+3+4+5
		},
		{
			name:     "Decreasing ratings",
			ratings:  []int{5, 4, 3, 2, 1},
			expected: 15, // 5+4+3+2+1
		},
		{
			name:     "Valley shape",
			ratings:  []int{3, 2, 1, 2, 3},
			expected: 11, // 3+2+1+2+3
		},
		{
			name:     "Peak shape",
			ratings:  []int{1, 2, 3, 2, 1},
			expected: 9, // 1+2+3+2+1
		},
		{
			name:     "Two children same rating",
			ratings:  []int{1, 1},
			expected: 2,
		},
		{
			name:     "Two children increasing",
			ratings:  []int{1, 2},
			expected: 3, // 1+2
		},
		{
			name:     "Two children decreasing",
			ratings:  []int{2, 1},
			expected: 3, // 2+1
		},
		{
			name:     "Complex pattern",
			ratings:  []int{1, 3, 2, 2, 1},
			expected: 7, // 1+2+1+2+1
		},
		{
			name:     "Multiple peaks and valleys",
			ratings:  []int{1, 2, 3, 1, 0, 2, 3, 2, 1},
			expected: 17,
		},
	}

	for _, tc := range testCases {
		result := candy(tc.ratings)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Ratings:  %v\n", tc.ratings)
			fmt.Printf("  Expected: %d\n", tc.expected)
			fmt.Printf("  Got:      %d\n", result)
		}
	}
}
