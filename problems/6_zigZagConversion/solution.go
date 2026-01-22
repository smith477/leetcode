package main

import "fmt"

/*
=== LeetCode Problem 6: Zigzag Conversion ===

Problem Description:
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows:
P   A   H   N
A P L S I I G
Y   I   R

Then read line by line: "PAHNAPLSIIGYIR"

Approach:
---------
We use a row-by-row simulation approach:

1. Create slices for each row to collect characters
2. Use a currentRow variable to track which row we're on
3. Use a goingDown flag to track direction (down or up)
4. For each character in the string:
   - Add it to the current row
   - If we're at the top (row 0) or bottom (row numRows-1), flip direction
   - Move to next row based on direction
5. Concatenate all rows to get the result

Example walkthrough with "PAYPALISHIRING", numRows = 3:
- P goes to row 0 (going down)
- A goes to row 1 (going down)
- Y goes to row 2 (hit bottom, flip to going up)
- P goes to row 1 (going up)
- A goes to row 0 (hit top, flip to going down)
- L goes to row 1 (going down)
- And so on...

Result:
Row 0: P   A   H   N     = "PAHN"
Row 1: A P L S I I G     = "APLSII G"
Row 2: Y   I   R         = "YIR"
Final: "PAHNAPLSIIGYIR"

Time Complexity: O(n) where n is the length of the string - we visit each character once
Space Complexity: O(n) for storing characters in rows

Edge Cases:
- numRows = 1: return original string (no zigzag)
- len(s) <= numRows: return original string (not enough chars for zigzag)
*/

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	rows := make([][]byte, numRows)

	currentRow := 0
	goingDown := false

	for i := 0; i < len(s); i++ {
		rows[currentRow] = append(rows[currentRow], s[i])

		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	result := make([]byte, 0, len(s))
	for _, row := range rows {
		result = append(result, row...)
	}

	return string(result)
}

func main() {
	fmt.Println("=== LeetCode Problem 6: Zigzag Conversion ===\n")

	// Test Case 1
	testCase1 := "PAYPALISHIRING"
	numRows1 := 3
	result1 := convert(testCase1, numRows1)
	fmt.Printf("Test 1: s = %q, numRows = %d\n", testCase1, numRows1)
	fmt.Printf("Output: %q\n", result1)
	fmt.Printf("Expected: \"PAHNAPLSIIGYIR\"\n")
	fmt.Printf("Passed: %v\n\n", result1 == "PAHNAPLSIIGYIR")

	// Test Case 2
	testCase2 := "PAYPALISHIRING"
	numRows2 := 4
	result2 := convert(testCase2, numRows2)
	fmt.Printf("Test 2: s = %q, numRows = %d\n", testCase2, numRows2)
	fmt.Printf("Output: %q\n", result2)
	fmt.Printf("Expected: \"PINALSIGYAHRPI\"\n")
	fmt.Printf("Passed: %v\n\n", result2 == "PINALSIGYAHRPI")

	// Test Case 3: Edge case with numRows = 1
	testCase3 := "ABCDEFGH"
	numRows3 := 1
	result3 := convert(testCase3, numRows3)
	fmt.Printf("Test 3 (Edge): s = %q, numRows = %d\n", testCase3, numRows3)
	fmt.Printf("Output: %q\n", result3)
	fmt.Printf("Expected: \"ABCDEFGH\"\n")
	fmt.Printf("Passed: %v\n\n", result3 == "ABCDEFGH")

	// Visual demonstration
	fmt.Println("=== Visual Demonstration ===")
	fmt.Println("Input: \"PAYPALISHIRING\", numRows = 3")
	fmt.Println("\nZigzag pattern:")
	fmt.Println("P   A   H   N")
	fmt.Println("A P L S I I G")
	fmt.Println("Y   I   R")
	fmt.Printf("\nReading row by row: %s\n", convert("PAYPALISHIRING", 3))
}
