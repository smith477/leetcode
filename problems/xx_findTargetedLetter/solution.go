package main

import "fmt"

type inputType struct {
	letters []byte
	target  byte
}

type TestCase struct {
	input  inputType
	output byte
}

func binarySearch(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1

	for left <= right {
		mid := left + (right-left)/2 // Move this INSIDE the loop

		if letters[mid] == target {
			return letters[mid]
		} else if letters[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}

func main() {

	testCases := []TestCase{
		// Empty array
		{
			input: inputType{
				letters: []byte{},
				target:  'a',
			},
			output: 0,
		},
		// Single element - found
		{
			input: inputType{
				letters: []byte{'m'},
				target:  'm',
			},
			output: 'm',
		},
		// Single element - not found
		{
			input: inputType{
				letters: []byte{'m'},
				target:  'z',
			},
			output: 0,
		},
		// Two elements - find first
		{
			input: inputType{
				letters: []byte{'a', 'z'},
				target:  'a',
			},
			output: 'a',
		},
		// Two elements - find second
		{
			input: inputType{
				letters: []byte{'a', 'z'},
				target:  'z',
			},
			output: 'z',
		},
		// Target before all elements
		{
			input: inputType{
				letters: []byte{'d', 'f', 'h', 'j', 'l', 'n', 'p'},
				target:  'a',
			},
			output: 0,
		},
		// Target after all elements
		{
			input: inputType{
				letters: []byte{'d', 'f', 'h', 'j', 'l', 'n', 'p'},
				target:  'z',
			},
			output: 0,
		},
		// Target in middle (exists)
		{
			input: inputType{
				letters: []byte{'b', 'd', 'f', 'h', 'j', 'l', 'n', 'p', 'r'},
				target:  'j',
			},
			output: 'j',
		},
		// Target between elements (doesn't exist)
		{
			input: inputType{
				letters: []byte{'b', 'd', 'f', 'h', 'j', 'l', 'n', 'p', 'r'},
				target:  'k',
			},
			output: 0,
		},
		// Consecutive letters
		{
			input: inputType{
				letters: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
				target:  'e',
			},
			output: 'e',
		},
		// Large gaps between letters
		{
			input: inputType{
				letters: []byte{'a', 'e', 'i', 'o', 'u', 'y'},
				target:  'o',
			},
			output: 'o',
		},
		// All same letters (edge case)
		{
			input: inputType{
				letters: []byte{'k', 'k', 'k', 'k', 'k'},
				target:  'k',
			},
			output: 'k',
		},
		// Alphabet endpoints
		{
			input: inputType{
				letters: []byte{'a', 'b', 'x', 'y', 'z'},
				target:  'z',
			},
			output: 'z',
		},
	}

	for i, tc := range testCases {
		result := binarySearch(tc.input.letters, tc.input.target)

		// Format expected and result for display
		expected := "empty"
		if tc.output != 0 {
			expected = fmt.Sprintf("'%c'", tc.output)
		}

		got := "empty"
		if result != 0 {
			got = fmt.Sprintf("'%c'", result)
		}

		fmt.Printf("Test %2d: Expected %-7s Got %s\n", i+1, expected, got)
	}
}
