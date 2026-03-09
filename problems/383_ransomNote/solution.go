package main

// Hash Table (Array) - Ransom Note
// Time complexity O(m + n), Space complexity O(1) - fixed 26-size array
// Create a frequency array of 26 elements for lowercase letters a-z
// First pass: count all characters in magazine by incrementing charCount[char-'a']
// Second pass: for each character in ransomNote, decrement its count
//   - If count goes negative, that character is not available enough times
//   - Return false immediately when this happens
// If we successfully process all characters in ransomNote, return true
func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	// Count frequency of each character in magazine
	charCount := [26]int{}
	for _, char := range magazine {
		charCount[char-'a']++
	}

	// Check if ransomNote can be constructed
	for _, char := range ransomNote {
		charCount[char-'a']--
		if charCount[char-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {
	testCases := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},                     // Different characters
		{"aa", "ab", false},                   // Not enough 'a's
		{"aa", "aab", true},                   // Exact match with extra
		{"", "abc", true},                     // Empty ransom note
		{"abc", "aabbcc", true},               // All characters available
		{"abc", "", false},                    // Empty magazine
		{"", "", true},                        // Both empty
		{"aab", "baa", true},                  // Order doesn't matter
		{"fffbfg", "effjfggbffjdgbjjhhdegh", true}, // LeetCode example
		{"bg", "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj", true}, // Long magazine
		{"aabbcc", "abc", false},              // Multiple duplicates needed
		{"aaa", "aaa", true},                  // Exact same
		{"aaab", "aaa", false},                // One character short
		{"leetcode", "letecode", true},        // Has 4 e's, needs 3
		{"leetcode", "leetcodeee", true},      // Extra characters in magazine
		{"xyz", "abcdefghijklmnopqrstuvwxyz", true}, // All alphabet
		{"aaaaaa", "aaaaaab", true},           // Six a's available
	}

	for i, tc := range testCases {
		result := canConstruct(tc.ransomNote, tc.magazine)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		println("Test", i+1, "- Expected:", tc.expected, "| Result:", result, "|", status)
	}
}
