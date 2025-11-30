package main

import (
	"fmt"
	"strings"
)

// Split words manually, reverse array, join with single space
// Time: O(n), Space: O(n)
// Skip spaces, extract words into array, reverse array in-place, join
func reverseWords(s string) string {
	words := []string{}
	i := 0
	n := len(s)

	for i < n {
		// Skip spaces
		for i < n && s[i] == ' ' {
			i++
		}

		// Extract word
		if i < n {
			start := i
			for i < n && s[i] != ' ' {
				i++
			}
			words = append(words, s[start:i])
		}
	}

	// Reverse words array
	left, right := 0, len(words)-1
	reversed(words, left, right)

	return strings.Join(words, " ")
}

func reversed(s []string, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// Use strings.Fields to split, then reverse
// Time: O(n), Space: O(n)
// func reverseWords(s string) string {
// 	words := strings.Fields(s)  // Splits on whitespace, removes empty strings

// 	left, right := 0, len(words)-1
// 	for left < right {
// 		words[left], words[right] = words[right], words[left]
// 		left++
// 		right--
// 	}

// 	return strings.Join(words, " ")
// }

func main() {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1",
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			name:     "Example 2",
			input:    "  hello world  ",
			expected: "world hello",
		},
		{
			name:     "Example 3",
			input:    "a good   example",
			expected: "example good a",
		},
		{
			name:     "Single word",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "Single word with spaces",
			input:    "  hello  ",
			expected: "hello",
		},
		{
			name:     "Two words",
			input:    "hello world",
			expected: "world hello",
		},
		{
			name:     "Two words with multiple spaces",
			input:    "hello    world",
			expected: "world hello",
		},
		{
			name:     "Leading spaces",
			input:    "   hello world",
			expected: "world hello",
		},
		{
			name:     "Trailing spaces",
			input:    "hello world   ",
			expected: "world hello",
		},
		{
			name:     "Leading and trailing spaces",
			input:    "   hello world   ",
			expected: "world hello",
		},
		{
			name:     "Multiple spaces between words",
			input:    "a  b   c    d",
			expected: "d c b a",
		},
		{
			name:     "Many spaces everywhere",
			input:    "  Bob    Loves  Alice   ",
			expected: "Alice Loves Bob",
		},
		{
			name:     "Single character words",
			input:    "a b c",
			expected: "c b a",
		},
		{
			name:     "Mixed case",
			input:    "Hello World",
			expected: "World Hello",
		},
		{
			name:     "With digits",
			input:    "test 123 hello",
			expected: "hello 123 test",
		},
		{
			name:     "Long words",
			input:    "programming is awesome",
			expected: "awesome is programming",
		},
		{
			name:     "Three words with varied spacing",
			input:    "  one   two  three  ",
			expected: "three two one",
		},
	}

	for _, tc := range testCases {
		result := reverseWords(tc.input)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Input:    %q\n", tc.input)
			fmt.Printf("  Expected: %q\n", tc.expected)
			fmt.Printf("  Got:      %q\n", result)
		}
	}
}
