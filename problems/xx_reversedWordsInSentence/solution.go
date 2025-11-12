package main

import (
	"fmt"
	"strings"
)

/*
Here's a more readable version of the problem:

---

## Problem: Fix Reversed Words in a Sentence

You receive a sentence where something went wrong during transmission:
- Each word has its letters **reversed**
- Words are separated by **one or more spaces** (sometimes too many!)
- There might be **leading or trailing spaces**

Your task is to fix the sentence by:
1. Reversing the letters in each word back to normal
2. Normalizing all spacing to single spaces
3. Removing any leading or trailing spaces

### Examples

**Example 1:**
```
Input:  "sihT   si   a  tset"
Output: "This is a test"
```

**Example 2:**
```
Input:  "  olleH   dlroW "
Output: "Hello World"
```

**Example 3:**
```
Input:  "eno   owt  eerht"
Output: "one two three"
```

### Constraints
- `0 ≤ s.length ≤ 100,000`
- String contains only ASCII letters and spaces
- **Must run in O(n) time**

### Bonus Challenge
If the string can contain Unicode characters (emojis, accented letters, etc.), you need to be careful to reverse by **user-perceived characters** (grapheme clusters), not just bytes.

*/

func fixReversedSentence(s string) string {
	if len(s) == 0 {
		return ""
	}

	words := []string{}
	currentIndex := 0

	for currentIndex < len(s) {
		for currentIndex < len(s) && s[currentIndex] == ' ' {
			currentIndex++
		}

		if currentIndex >= len(s) {
			break
		}

		wordStart := currentIndex
		for currentIndex < len(s) && s[currentIndex] != ' ' {
			currentIndex++
		}

		word := s[wordStart:currentIndex]
		reversed := reverseString(word)
		words = append(words, reversed)
	}

	return strings.Join(words, " ")
}

func reverseString(s string) string {
	runes := []rune(s)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

func main() {

	testCases := []struct {
		input  string
		output string
	}{
		{"sihT   si   a  tset", "This is a test"},
		{"  olleH   dlroW ", "Hello World"},
		{"eno owt eerht", "one two three"},
	}

	for index, testCase := range testCases {
		fmt.Printf("Test %d - Expected: %s, Got: %s\n", index, testCase.output, fixReversedSentence(testCase.input))
	}
}
