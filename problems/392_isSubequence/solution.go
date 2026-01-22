package main

import "fmt"

// Two-pointer subsequence matching
// Time: O(n), Space: O(1) where n is length of t
// Iterate through t, advance pointer in s when characters match
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sIdx, tIdx := 0, 0

	for sIdx < len(s) && tIdx < len(t) {
		if s[sIdx] == t[tIdx] {
			sIdx++
		}
		tIdx++
	}

	return sIdx == len(s)
}

func main() {
	// Test cases
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"aec", "abcde", false},
		{"axc", "ahbgdc", false},
		{"", "ahbgdc", true},
		{"b", "abc", true},
	}

	for i, tc := range testCases {
		result := isSubsequence(tc.s, tc.t)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("%s Test %d: isSubsequence(%q, %q) = %v (expected %v)\n",
			status, i+1, tc.s, tc.t, result, tc.expected)
	}
}
