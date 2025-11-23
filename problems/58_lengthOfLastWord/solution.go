package main

import "fmt"

// Reverse traversal to find last word
// Time: O(n), Space: O(1)
// Skip trailing spaces, then count characters until next space or start
func lengthOfLastWord(s string) int {

	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}

	length := 0
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}

func main() {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Example 1",
			input:    "Hello World",
			expected: 5,
		},
		{
			name:     "Example 2",
			input:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "Example 3",
			input:    "luffy is still joyboy",
			expected: 6,
		},
		{
			name:     "Single word",
			input:    "hello",
			expected: 5,
		},
		{
			name:     "Single word with leading spaces",
			input:    "   hello",
			expected: 5,
		},
		{
			name:     "Single word with trailing spaces",
			input:    "hello   ",
			expected: 5,
		},
		{
			name:     "Single word with spaces both sides",
			input:    "   hello   ",
			expected: 5,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: 1,
		},
		{
			name:     "Two words",
			input:    "hello world",
			expected: 5,
		},
		{
			name:     "Multiple spaces between words",
			input:    "hello     world",
			expected: 5,
		},
		{
			name:     "Long last word",
			input:    "a supercalifragilisticexpialidocious",
			expected: 34,
		},
		{
			name:     "Many words",
			input:    "the quick brown fox jumps",
			expected: 5,
		},
		{
			name:     "Trailing spaces after last word",
			input:    "test word     ",
			expected: 4,
		},
		{
			name:     "Many trailing spaces",
			input:    "a b          ",
			expected: 1,
		},
	}

	for _, tc := range testCases {
		result := lengthOfLastWord(tc.input)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Input:    %q\n", tc.input)
			fmt.Printf("  Expected: %d\n", tc.expected)
			fmt.Printf("  Got:      %d\n", result)
		}
	}
}
