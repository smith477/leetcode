package main

import "fmt"

func printMatrix(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		fmt.Print("[")
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%2d", matrix[row][col])
			if col < len(matrix[row])-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")
	}
}

// Matrix manipulation: transpose then reverse rows
// Time: O(n²), Space: O(1)
// Transpose matrix (swap across diagonal), then reverse each row to rotate 90° clockwise
func rotate(matrix [][]int) {
	n := len(matrix)

	// Transpose: swap across diagonal
	for row := 0; row < n; row++ {
		for col := row + 1; col < n; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	// Reverse each row
	for row := 0; row < n; row++ {
		left, right := 0, n-1
		for left < right {
			matrix[row][left], matrix[row][right] = matrix[row][right], matrix[row][left]
			left++
			right--
		}
	}
}

func main() {
	// Test case 1: 3x3 matrix
	fmt.Println("=== Test 1: 3x3 matrix ===")
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Before rotation:")
	printMatrix(matrix1)
	fmt.Println()

	rotate(matrix1)

	fmt.Println("\nAfter rotation:")
	printMatrix(matrix1)
	fmt.Println("Expected:")
	fmt.Println("[ 7  4  1]")
	fmt.Println("[ 8  5  2]")
	fmt.Println("[ 9  6  3]")
	fmt.Println()

	// Test case 2: 4x4 matrix
	fmt.Println("=== Test 2: 4x4 matrix ===")
	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println("Before rotation:")
	printMatrix(matrix2)
	fmt.Println()

	rotate(matrix2)

	fmt.Println("\nAfter rotation:")
	printMatrix(matrix2)
	fmt.Println("Expected:")
	fmt.Println("[13  9  5  1]")
	fmt.Println("[14 10  6  2]")
	fmt.Println("[15 11  7  3]")
	fmt.Println("[16 12  8  4]")
	fmt.Println()

	// Test case 3: 1x1 matrix
	fmt.Println("=== Test 3: 1x1 matrix ===")
	matrix3 := [][]int{
		{1},
	}
	fmt.Println("Before rotation:")
	printMatrix(matrix3)
	fmt.Println()

	rotate(matrix3)

	fmt.Println("\nAfter rotation:")
	printMatrix(matrix3)
	fmt.Println("Expected:")
	fmt.Println("[ 1]")
	fmt.Println()

	// Test case 4: 2x2 matrix
	fmt.Println("=== Test 4: 2x2 matrix ===")
	matrix4 := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("Before rotation:")
	printMatrix(matrix4)
	fmt.Println()

	rotate(matrix4)

	fmt.Println("\nAfter rotation:")
	printMatrix(matrix4)
	fmt.Println("Expected:")
	fmt.Println("[ 3  1]")
	fmt.Println("[ 4  2]")
}
