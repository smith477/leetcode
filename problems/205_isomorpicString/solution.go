package main

import "fmt"

//func isIsomorphic(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	var mapS [256]int
//	var mapT [256]int
//
//	for i := 0; i < len(s); i++ {
//		charS := s[i]
//		charT := t[i]
//
//		if mapS[charS] != mapT[charT] {
//			return false
//		}
//
//		mapS[charS] = i + 1
//		mapT[charT] = i + 1
//	}
//
//	return true
//}

func main() {
	tests := []struct {
		s, t     string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"f11", "b23", false},
		{"ab", "aa", false},     // additional test
		{"badc", "baba", false}, // additional test
	}

	for _, test := range tests {
		result := isIsomorphic(test.s, test.t)
		fmt.Printf("isIsomorphic(%q, %q) = %v, expected %v → %v\n",
			test.s, test.t, result, test.expected, result == test.expected)
	}
}
