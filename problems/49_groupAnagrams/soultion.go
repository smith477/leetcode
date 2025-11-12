package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	anagramMap := make(map[string][]string)

	for _, word := range strs {
		key := sortedString(word)

		anagramMap[key] = append(anagramMap[key], word)
	}

	fmt.Printf("Test: %v\n", anagramMap)

	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func sortedString(s string) string {
	chars := []rune(s)

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}

func main() {
	// Test Case 1
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("Input: %v\n", strs1)
	fmt.Printf("Output: %v\n\n", groupAnagrams(strs1))

	// Test Case 2
	strs2 := []string{""}
	fmt.Printf("Input: %v\n", strs2)
	fmt.Printf("Output: %v\n\n", groupAnagrams(strs2))

	// Test Case 3
	strs3 := []string{"a"}
	fmt.Printf("Input: %v\n", strs3)
	fmt.Printf("Output: %v\n\n", groupAnagrams(strs3))
}
