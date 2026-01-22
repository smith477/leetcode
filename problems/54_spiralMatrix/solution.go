package main

import "fmt"

// Layer-by-layer traversal using four boundaries
// Time: O(m*n), Space: O(1) - excluding result array
// Traverse right→down→left→up, shrinking boundaries after each direction
func spiralOrder(matrix [][]int) []int {
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	result := []int{}

	for top <= bottom && left <= right {
		// Traverse right across the top row
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++ // Move top boundary down

		// Traverse down the right column
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right-- // Move right boundary left

		// Traverse left across the bottom row (if row still exists)
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom-- // Move bottom boundary up
		}

		// Traverse up the left column (if column still exists)
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++ // Move left boundary right
		}
	}
	return result
}

// 00, 01, 02, 12, 22, 21, 20, 10, 11
// 00, 01, 02, 03, 13, 23, 22, 21, 20, 10, 11, 12
func main() {
	// Test case 1: 3x3 matrix
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Test 1 - 3x3 matrix:")
	fmt.Println("Input:", matrix1)
	fmt.Println("Output:", spiralOrder(matrix1))
	fmt.Println("Expected: [1 2 3 6 9 8 7 4 5]")
	fmt.Println()

	// Test case 2: 3x4 matrix
	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println("Test 2 - 3x4 matrix:")
	fmt.Println("Input:", matrix2)
	fmt.Println("Output:", spiralOrder(matrix2))
	fmt.Println("Expected: [1 2 3 4 8 12 11 10 9 5 6 7]")
	fmt.Println()

	// Test case 3: 1x4 matrix (single row)
	matrix3 := [][]int{
		{1, 2, 3, 4},
	}
	fmt.Println("Test 3 - 1x4 matrix (single row):")
	fmt.Println("Input:", matrix3)
	fmt.Println("Output:", spiralOrder(matrix3))
	fmt.Println("Expected: [1 2 3 4]")
	fmt.Println()

	// Test case 4: 4x1 matrix (single column)
	matrix4 := [][]int{
		{1},
		{2},
		{3},
		{4},
	}
	fmt.Println("Test 4 - 4x1 matrix (single column):")
	fmt.Println("Input:", matrix4)
	fmt.Println("Output:", spiralOrder(matrix4))
	fmt.Println("Expected: [1 2 3 4]")
	fmt.Println()

	// Test case 5: 1x1 matrix (single element)
	matrix5 := [][]int{
		{5},
	}
	fmt.Println("Test 5 - 1x1 matrix (single element):")
	fmt.Println("Input:", matrix5)
	fmt.Println("Output:", spiralOrder(matrix5))
	fmt.Println("Expected: [5]")
}
