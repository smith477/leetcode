package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	n := len(haystack)
	m := len(needle)

	if m > n {
		return -1
	}

	for i := 0; i <= n-m; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}

	return -1
}

func main() {
	testCases := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "Example 1",
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			name:     "Example 2",
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
		{
			name:     "Needle at beginning",
			haystack: "hello",
			needle:   "he",
			expected: 0,
		},
		{
			name:     "Needle at end",
			haystack: "hello",
			needle:   "lo",
			expected: 3,
		},
		{
			name:     "Needle in middle",
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
		{
			name:     "Needle is entire haystack",
			haystack: "abc",
			needle:   "abc",
			expected: 0,
		},
		{
			name:     "Single character needle found",
			haystack: "abc",
			needle:   "a",
			expected: 0,
		},
		{
			name:     "Single character needle not found",
			haystack: "abc",
			needle:   "d",
			expected: -1,
		},
		{
			name:     "Needle longer than haystack",
			haystack: "ab",
			needle:   "abc",
			expected: -1,
		},
		{
			name:     "Multiple occurrences, return first",
			haystack: "aaaaa",
			needle:   "aa",
			expected: 0,
		},
		{
			name:     "Needle appears multiple times",
			haystack: "mississippi",
			needle:   "issip",
			expected: 4,
		},
		{
			name:     "Similar but not matching",
			haystack: "aaa",
			needle:   "aaaa",
			expected: -1,
		},
		{
			name:     "Partial match then full match",
			haystack: "ababcabc",
			needle:   "abc",
			expected: 2,
		},
		{
			name:     "Repeating pattern",
			haystack: "abababab",
			needle:   "abab",
			expected: 0,
		},
		{
			name:     "No match - completely different",
			haystack: "hello",
			needle:   "world",
			expected: -1,
		},
		{
			name:     "Empty needle",
			haystack: "hello",
			needle:   "",
			expected: 0,
		},
		{
			name:     "Both empty",
			haystack: "",
			needle:   "",
			expected: 0,
		},
		{
			name:     "Haystack empty, needle not",
			haystack: "",
			needle:   "a",
			expected: -1,
		},
	}

	for _, tc := range testCases {
		result := strStr(tc.haystack, tc.needle)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Haystack: %q\n", tc.haystack)
			fmt.Printf("  Needle:   %q\n", tc.needle)
			fmt.Printf("  Expected: %d\n", tc.expected)
			fmt.Printf("  Got:      %d\n", result)
		}
	}
}
