package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	sFields := strings.Fields(s)

	if len(pattern) != len(sFields) {
		return false
	}

	wordMap := map[byte]string{}
	charMap := map[string]byte{}

	for i := 0; i < len(pattern); i++ {
		patternChar := pattern[i]
		word := sFields[i]

		if foundWord, found := wordMap[patternChar]; found {
			if word != foundWord {
				return false
			}
		}

		if foundChar, found := charMap[word]; found {
			if patternChar != foundChar {
				return false
			}
		}

		wordMap[patternChar] = word
		charMap[word] = patternChar
	}

	return true
}

func main() {
	testCases := []struct {
		pattern string
		s       string
		want    bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
		{"abc", "dog cat fish", true},
	}

	for i, tc := range testCases {
		got := wordPattern(tc.pattern, tc.s)
		if got == tc.want {
			fmt.Printf("Test %d: PASS\n", i+1)
		} else {
			fmt.Printf("Test %d: FAIL (pattern=%q, s=%q, want=%v, got=%v)\n", i+1, tc.pattern, tc.s, tc.want, got)
		}
	}
}
