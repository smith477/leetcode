package main

import "fmt"

// Two-pass boundary approach
// Time: O(n), Space: O(1)
// Pass 1: Move left-to-right, find boundaries >= current height and calculate water
// When no valid boundary found, mark position and switch to pass 2
// Pass 2: Move right-to-left from end back to where pass 1 stopped
// This ensures all trapped water is counted even when heights decrease towards the end
// func trap(height []int) int {
// 	if len(height) <= 2 {
// 		return 0
// 	}

// 	n := len(height)
// 	totalWater := 0

// 	// Pass 1: Left to right until we can't find a higher/equal boundary
// 	left := 0
// 	for left < n-1 {
// 		// Find next boundary that's >= current height
// 		right := left + 1
// 		for right < n && height[right] < height[left] {
// 			right++
// 		}

// 		// If found a valid boundary, calculate water
// 		if right < n {
// 			for i := left + 1; i < right; i++ {
// 				totalWater += height[left] - height[i]
// 			}
// 			left = right
// 		} else {
// 			// No valid boundary found, stop left-to-right pass
// 			break
// 		}
// 	}

// 	// Pass 2: Right to left from the end back to where we stopped
// 	right := n - 1
// 	for right > left {
// 		nextLeft := right - 1
// 		for nextLeft >= left && height[nextLeft] < height[right] {
// 			nextLeft--
// 		}

// 		// Calculate water between boundaries
// 		if nextLeft >= left { // Changed from > to >=
// 			for i := nextLeft + 1; i < right; i++ {
// 				totalWater += height[right] - height[i]
// 			}
// 			right = nextLeft
// 		} else {
// 			break
// 		}
// 	}

// 	return totalWater
// }

// Two-pointer approach
// Time: O(n), Space: O(1)
// Move pointers inward from both ends, tracking max heights seen from each side
// Water trapped at each position is limited by the smaller of the two max heights
// Process the side with smaller height first to ensure valid water calculation
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}

func main() {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			name:     "Example 2",
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
		{
			name:     "Empty array",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			height:   []int{5},
			expected: 0,
		},
		{
			name:     "Two elements",
			height:   []int{3, 5},
			expected: 0,
		},
		{
			name:     "All same height",
			height:   []int{3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "Increasing heights",
			height:   []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "Decreasing heights",
			height:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Simple valley",
			height:   []int{3, 0, 3},
			expected: 3,
		},
		{
			name:     "Deep valley",
			height:   []int{5, 0, 0, 0, 5},
			expected: 15,
		},
		{
			name:     "Multiple valleys",
			height:   []int{3, 0, 2, 0, 4},
			expected: 7,
		},
		{
			name:     "Peak in middle",
			height:   []int{2, 1, 3, 1, 2},
			expected: 2,
		},
		{
			name:     "No water trapped",
			height:   []int{1, 2, 1, 2, 1},
			expected: 1,
		},
		{
			name:     "All zeros",
			height:   []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Zeros at ends",
			height:   []int{0, 1, 0, 1, 0},
			expected: 1,
		},
		{
			name:     "Step pattern",
			height:   []int{5, 4, 3, 2, 1, 2, 3, 4, 5},
			expected: 16,
		},
		{
			name:     "Large numbers",
			height:   []int{100, 0, 0, 0, 100},
			expected: 300,
		},
		{
			name:     "Example 3",
			height:   []int{4, 2, 3},
			expected: 1,
		},
		{
			name:     "Example 4",
			height:   []int{0, 0, 0},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		result := trap(tc.height)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Height:   %v\n", tc.height)
			fmt.Printf("  Expected: %d\n", tc.expected)
			fmt.Printf("  Got:      %d\n", result)
		}
	}
}
