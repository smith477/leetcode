package main

import "fmt"

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	charMap := make(map[rune]int)

	for _, char := range s {
		charMap[char] += 1
	}

	for _, char := range t {
		charMap[char]--

		if charMap[char] < 0 {
			return false
		}
	}

	return true
}

func main() {
	// Test cases
	testCases := []struct {
		s        string
		t        string
		expected bool
		desc     string
	}{
		{"anagram", "nagaram", true, "Example 1: Basic anagram"},
		{"rat", "car", false, "Example 2: Not an anagram"},
		{"", "", true, "Empty strings"},
		{"a", "a", true, "Single character match"},
		{"a", "b", false, "Single character mismatch"},
		{"listen", "silent", true, "Common anagram"},
		{"hello", "world", false, "Different lengths"},
		{"aabbcc", "abcabc", true, "Multiple occurrences"},
		{"abc", "abcd", false, "Different lengths"},
		{"café", "éfac", true, "Unicode characters (accented)"},
		{"こんにちは", "にちはこん", true, "Unicode characters (Japanese)"},
		{"😀😁😂", "😂😁😀", true, "Unicode characters (emojis)"},
	}

	fmt.Println("Running Valid Anagram Tests:")

	passCount := 0
	failCount := 0

	for i, tc := range testCases {
		result := isAnagram(tc.s, tc.t)
		status := "✓ PASS"
		if result != tc.expected {
			status = "✗ FAIL"
			failCount++
		} else {
			passCount++
		}

		fmt.Printf("Test %2d: %s\n", i+1, status)
		fmt.Printf("  Description: %s\n", tc.desc)
		fmt.Printf("  Input: s = \"%s\", t = \"%s\"\n", tc.s, tc.t)
		fmt.Printf("  Expected: %t, Got: %t\n", tc.expected, result)
		fmt.Println()
	}

	fmt.Printf("Results: %d passed, %d failed out of %d tests\n",
		passCount, failCount, len(testCases))
}
