package main

import "fmt"

func intToRoman(num int) string {

	values := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result string
	for _, v := range values {
		if num == 0 {
			break
		}

		for num >= v.value {
			result += v.symbol
			num -= v.value
		}
	}

	return result
}

func main() {
	testCases := []struct {
		name     string
		num      int
		expected string
	}{
		{
			name:     "Example 1",
			num:      3749,
			expected: "MMMDCCXLIX",
		},
		{
			name:     "Example 2",
			num:      58,
			expected: "LVIII",
		},
		{
			name:     "Example 3",
			num:      1994,
			expected: "MCMXCIV",
		},
		{
			name:     "Minimum value",
			num:      1,
			expected: "I",
		},
		{
			name:     "Maximum value",
			num:      3999,
			expected: "MMMCMXCIX",
		},
		{
			name:     "Single digit - 4",
			num:      4,
			expected: "IV",
		},
		{
			name:     "Single digit - 5",
			num:      5,
			expected: "V",
		},
		{
			name:     "Single digit - 9",
			num:      9,
			expected: "IX",
		},
		{
			name:     "Tens - 10",
			num:      10,
			expected: "X",
		},
		{
			name:     "Tens - 40",
			num:      40,
			expected: "XL",
		},
		{
			name:     "Tens - 50",
			num:      50,
			expected: "L",
		},
		{
			name:     "Tens - 90",
			num:      90,
			expected: "XC",
		},
		{
			name:     "Hundreds - 100",
			num:      100,
			expected: "C",
		},
		{
			name:     "Hundreds - 400",
			num:      400,
			expected: "CD",
		},
		{
			name:     "Hundreds - 500",
			num:      500,
			expected: "D",
		},
		{
			name:     "Hundreds - 900",
			num:      900,
			expected: "CM",
		},
		{
			name:     "Thousands - 1000",
			num:      1000,
			expected: "M",
		},
		{
			name:     "Thousands - 3000",
			num:      3000,
			expected: "MMM",
		},
		{
			name:     "Complex - 27",
			num:      27,
			expected: "XXVII",
		},
		{
			name:     "Complex - 49",
			num:      49,
			expected: "XLIX",
		},
		{
			name:     "Complex - 444",
			num:      444,
			expected: "CDXLIV",
		},
		{
			name:     "Complex - 1904",
			num:      1904,
			expected: "MCMIV",
		},
		{
			name:     "All same symbols - 3",
			num:      3,
			expected: "III",
		},
		{
			name:     "All same symbols - 30",
			num:      30,
			expected: "XXX",
		},
		{
			name:     "All same symbols - 300",
			num:      300,
			expected: "CCC",
		},
	}

	for _, tc := range testCases {
		result := intToRoman(tc.num)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Input:    %d\n", tc.num)
			fmt.Printf("  Expected: %s\n", tc.expected)
			fmt.Printf("  Got:      %s\n", result)
		}
	}
}
