package main

import (
	"fmt"
)

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

// func isAlphanumericRune(c rune) bool {
// 	return unicode.IsDigit(c) || unicode.IsLetter(c)
// }

func toLowerCaseIfNeeded(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func isPalindrome(s string) bool {

	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		leftChar := toLowerCaseIfNeeded(s[left])
		rightChar := toLowerCaseIfNeeded(s[right])

		if leftChar != rightChar {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	testCases := []struct {
		input  string
		output bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"!!!", true},
	}

	for index, testCase := range testCases {
		result := isPalindrome(testCase.input)
		fmt.Printf("Test %d - Expected: %v, Got: %v\n", index, testCase.output, result)
	}
}
