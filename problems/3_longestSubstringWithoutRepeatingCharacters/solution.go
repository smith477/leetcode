package main

import "fmt"

// Two pointer pattern
// Time complexity O(n), Space complexity O(n)
// Use hash map to remember bytes and their index
// One pointer is going forward one is staying at the beginning of the subarray
// Until repeating character is found right pointer progresses and with it max length increases
// When repeating character is found in hash map move left pointer on the positionIndex + 1 so new subarray can be found
func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	left := 0

	maxLen := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		if lastIndex, found := charMap[char]; found && lastIndex >= left {
			left = lastIndex + 1
		}

		charMap[char] = right

		currentLen := right - left + 1

		maxLen = max(currentLen, maxLen)
	}

	return maxLen
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"au", 2},
		{"dvdf", 3},
		{"abba", 2},
		{"tmmzuxt", 5},
	}

	fmt.Println("Running test cases for Longest Substring Without Repeating Characters...\n")

	for i, tc := range testCases {
		result := lengthOfLongestSubstring(tc.input)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("Test %d: %s\n", i+1, status)
		fmt.Printf("  Input:    \"%s\"\n", tc.input)
		fmt.Printf("  Expected: %d\n", tc.expected)
		fmt.Printf("  Got:      %d\n", result)
		fmt.Println()
	}
}
