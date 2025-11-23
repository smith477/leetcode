package main

import "fmt"

// Horizontal scanning approach
// Time: O(S) where S is sum of all characters, Space: O(1)
// Compare each string with current prefix, shrink prefix when mismatch found
// Stop early if prefix becomes empty
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, word := range strs[1:] {
		i := 0
		for i < len(prefix) && i < len(word) && prefix[i] == word[i] {
			i++
		}
		prefix = prefix[:i]

		if len(prefix) == 0 {
			return ""
		}
	}

	return prefix
}

func main() {
	testCases := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			name:     "Example 1",
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "Example 2",
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "Single string",
			strs:     []string{"hello"},
			expected: "hello",
		},
		{
			name:     "Two identical strings",
			strs:     []string{"test", "test"},
			expected: "test",
		},
		{
			name:     "All identical strings",
			strs:     []string{"abc", "abc", "abc"},
			expected: "abc",
		},
		{
			name:     "No common prefix",
			strs:     []string{"abc", "def", "ghi"},
			expected: "",
		},
		{
			name:     "Empty string in array",
			strs:     []string{"", "abc", "def"},
			expected: "",
		},
		{
			name:     "One empty, others with prefix",
			strs:     []string{"flower", "", "flow"},
			expected: "",
		},
		{
			name:     "Common prefix of one character",
			strs:     []string{"a", "ab", "abc"},
			expected: "a",
		},
		{
			name:     "First string is prefix of others",
			strs:     []string{"fl", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "Last string is shortest",
			strs:     []string{"flower", "flow", "fl"},
			expected: "fl",
		},
		{
			name:     "Different lengths, same prefix",
			strs:     []string{"prefix", "pre", "prepare"},
			expected: "pre",
		},
		{
			name:     "Long common prefix",
			strs:     []string{"international", "internet", "interval"},
			expected: "inter",
		},
		{
			name:     "Two strings, one empty",
			strs:     []string{"abc", ""},
			expected: "",
		},
		{
			name:     "Prefix ends at different positions",
			strs:     []string{"abcdef", "abcxyz", "abc"},
			expected: "abc",
		},
		{
			name:     "All single character, same",
			strs:     []string{"a", "a", "a"},
			expected: "a",
		},
		{
			name:     "All single character, different",
			strs:     []string{"a", "b", "c"},
			expected: "",
		},
	}

	for _, tc := range testCases {
		result := longestCommonPrefix(tc.strs)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Input:    %v\n", tc.strs)
			fmt.Printf("  Expected: %q\n", tc.expected)
			fmt.Printf("  Got:      %q\n", result)
		}
	}
}
