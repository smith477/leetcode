package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	n := len(s)

	for i := 0; i < n; i++ {
		current := romanMap[s[i]]

		if i < n-1 && romanMap[s[i+1]] > current {
			sum -= current
		} else {
			sum += current
		}
	}

	return sum
}

func main() {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Example 1",
			input:    "III",
			expected: 3,
		},
		{
			name:     "Example 2",
			input:    "LVIII",
			expected: 58,
		},
		{
			name:     "Example 3",
			input:    "MCMXCIV",
			expected: 1994,
		},
		{
			name:     "Single character - I",
			input:    "I",
			expected: 1,
		},
		{
			name:     "Single character - V",
			input:    "V",
			expected: 5,
		},
		{
			name:     "Single character - X",
			input:    "X",
			expected: 10,
		},
		{
			name:     "Single character - L",
			input:    "L",
			expected: 50,
		},
		{
			name:     "Single character - C",
			input:    "C",
			expected: 100,
		},
		{
			name:     "Single character - D",
			input:    "D",
			expected: 500,
		},
		{
			name:     "Single character - M",
			input:    "M",
			expected: 1000,
		},
		{
			name:     "Subtractive - IV",
			input:    "IV",
			expected: 4,
		},
		{
			name:     "Subtractive - IX",
			input:    "IX",
			expected: 9,
		},
		{
			name:     "Subtractive - XL",
			input:    "XL",
			expected: 40,
		},
		{
			name:     "Subtractive - XC",
			input:    "XC",
			expected: 90,
		},
		{
			name:     "Subtractive - CD",
			input:    "CD",
			expected: 400,
		},
		{
			name:     "Subtractive - CM",
			input:    "CM",
			expected: 900,
		},
		{
			name:     "Simple addition",
			input:    "VI",
			expected: 6,
		},
		{
			name:     "Simple addition - VII",
			input:    "VII",
			expected: 7,
		},
		{
			name:     "Complex - XXVII",
			input:    "XXVII",
			expected: 27,
		},
		{
			name:     "Complex - XLIX",
			input:    "XLIX",
			expected: 49,
		},
		{
			name:     "Large number - MMMCMXCIX",
			input:    "MMMCMXCIX",
			expected: 3999,
		},
		{
			name:     "Multiple subtractives - MCMIV",
			input:    "MCMIV",
			expected: 1904,
		},
	}

	for _, tc := range testCases {
		result := romanToInt(tc.input)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Input:    %s\n", tc.input)
			fmt.Printf("  Expected: %d\n", tc.expected)
			fmt.Printf("  Got:      %d\n", result)
		}
	}
}
