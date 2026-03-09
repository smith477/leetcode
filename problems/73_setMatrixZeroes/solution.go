package main

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	firstRow := false
	firstCol := false

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			firstCol = true
			break
		}
	}
	for j := 0; j < cols; j++ {
		if matrix[0][j] == 0 {
			firstRow = true
			break
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRow {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}

	if firstCol {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			print(val, " ")
		}
		println()
	}
}

func main() {
	// Test case 1: Standard case
	matrix1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	println("Test 1 - Before:")
	printMatrix(matrix1)
	setZeroes(matrix1)
	println("After:")
	printMatrix(matrix1)
	println("Expected: [[1,0,1],[0,0,0],[1,0,1]]")
	println()

	// Test case 2: Multiple zeros
	matrix2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	println("Test 2 - Before:")
	printMatrix(matrix2)
	setZeroes(matrix2)
	println("After:")
	printMatrix(matrix2)
	println("Expected: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]")
	println()

	// Test case 3: Single row
	matrix3 := [][]int{
		{1, 0, 3},
	}
	println("Test 3 - Before:")
	printMatrix(matrix3)
	setZeroes(matrix3)
	println("After:")
	printMatrix(matrix3)
	println("Expected: [[0,0,0]]")
	println()

	// Test case 4: No zeros
	matrix4 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	println("Test 4 - Before:")
	printMatrix(matrix4)
	setZeroes(matrix4)
	println("After:")
	printMatrix(matrix4)
	println("Expected: [[1,2,3],[4,5,6]]")
}
