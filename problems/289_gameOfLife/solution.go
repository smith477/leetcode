package main

func gameOfLife(board [][]int) {
	rows := len(board)
	cols := len(board[0])

	directions := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			live := 0
			for _, dir := range directions {
				newRow, newCol := i+dir[0], j+dir[1]

				if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
					continue
				}

				if board[newRow][newCol] == 1 || board[newRow][newCol] == 2 {
					live++
				}
			}

			if board[i][j] == 1 && (live < 2 || live > 3) {
				board[i][j] = 2
			} else if board[i][j] == 0 && live == 3 {
				board[i][j] = 3
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			board[i][j] = board[i][j] % 2
		}
	}
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			if cell == 1 {
				print("█ ")
			} else {
				print("· ")
			}
		}
		println()
	}
}

func copyBoard(board [][]int) [][]int {
	result := make([][]int, len(board))
	for i := range board {
		result[i] = make([]int, len(board[i]))
		copy(result[i], board[i])
	}
	return result
}

func main() {
	// Test case 1: Blinker (oscillator)
	board1 := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}
	expected1 := [][]int{
		{0, 0, 0},
		{1, 1, 1},
		{0, 0, 0},
	}
	println("Test 1 - Blinker:")
	println("Before:")
	printBoard(board1)
	gameOfLife(board1)
	println("After:")
	printBoard(board1)
	println("Expected:")
	printBoard(expected1)
	println()

	// Test case 2: Block (still life)
	board2 := [][]int{
		{1, 1},
		{1, 1},
	}
	expected2 := [][]int{
		{1, 1},
		{1, 1},
	}
	println("Test 2 - Block (should remain unchanged):")
	println("Before:")
	printBoard(board2)
	gameOfLife(board2)
	println("After:")
	printBoard(board2)
	println("Expected:")
	printBoard(expected2)
	println()

	// Test case 3: LeetCode example
	board3 := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	expected3 := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}
	println("Test 3 - LeetCode Example:")
	println("Before:")
	printBoard(board3)
	gameOfLife(board3)
	println("After:")
	printBoard(board3)
	println("Expected:")
	printBoard(expected3)
	println()

	// Test case 4: All dead cells
	board4 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	expected4 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	println("Test 4 - All dead (should remain dead):")
	println("Before:")
	printBoard(board4)
	gameOfLife(board4)
	println("After:")
	printBoard(board4)
	println("Expected:")
	printBoard(expected4)
}
