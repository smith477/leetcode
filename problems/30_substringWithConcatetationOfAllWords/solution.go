package main

import "fmt"

// Sliding window with word-length offsets
// Time: O(n * wordSize), Space: O(wordCount)
// Try each offset (0 to wordSize-1), use sliding window to match all words
func findSubstring(s string, words []string) []int {
	wordsMap := make(map[string]int)

	results := []int{}
	for _, word := range words {
		wordsMap[word]++
	}

	wordSize := len(words[0])
	wordCount := len(words)

	for offset := 0; offset < wordSize; offset++ {
		seen := make(map[string]int)
		left := offset
		matched := 0

		for right := offset; right+wordSize <= len(s); right += wordSize {
			word := s[right : right+wordSize]

			if _, found := wordsMap[word]; found {
				seen[word]++
				if seen[word] <= wordsMap[word] {
					matched++
				}
			}

			if (right-left)/wordSize >= wordCount {
				leftWord := s[left : left+wordSize]
				if _, found := wordsMap[leftWord]; found {
					if seen[leftWord] <= wordsMap[leftWord] {
						matched--
					}
					seen[leftWord]--
				}
				left += wordSize
			}

			if matched == wordCount {
				results = append(results, left)
			}
		}
	}

	return results
}

func main() {
	testCases := []struct {
		s        string
		words    []string
		expected []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
		{"lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}, []int{13}},
	}

	fmt.Println("Running test cases for Substring with Concatenation of All Words...\n")

	for i, tc := range testCases {
		result := findSubstring(tc.s, tc.words)
		passed := compareSlices(result, tc.expected)

		status := " PASS"
		if !passed {
			status = " FAIL"
		}

		fmt.Printf("Test %d: %s\n", i+1, status)
		fmt.Printf("  Input s:  \"%s\"\n", tc.s)
		fmt.Printf("  Words:    %v\n", tc.words)
		fmt.Printf("  Expected: %v\n", tc.expected)
		fmt.Printf("  Got:      %v\n", result)
		fmt.Println()
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	// Create maps to count occurrences
	aMap := make(map[int]int)
	bMap := make(map[int]int)

	for _, v := range a {
		aMap[v]++
	}

	for _, v := range b {
		bMap[v]++
	}

	// Compare maps
	for k, v := range aMap {
		if bMap[k] != v {
			return false
		}
	}

	return true
}
