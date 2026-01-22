package main

import "fmt"

// Hash set approach for rows, columns, and 3x3 boxes
// Time: O(1) - fixed 9x9 board, Space: O(1) - fixed size hash sets
// Track seen numbers in each row, column, and box using hash sets
func isValidSudoku(board [][]byte) bool {
	rows := [9]map[byte]bool{}
	cols := [9]map[byte]bool{}
	boxes := [9]map[byte]bool{}

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			val := board[row][col]

			if val == '.' {
				continue
			}

			boxIdx := (row/3)*3 + col/3

			if rows[row][val] || cols[col][val] || boxes[boxIdx][val] {
				return false
			}

			rows[row][val] = true
			cols[col][val] = true
			boxes[boxIdx][val] = true
		}
	}

	return true
}

//Using bits
//func isValidSudoku(board [][]byte) bool {
//    var rows, cols, boxes [9]int
//
//    for r := 0; r < 9; r++ {
//        for c := 0; c < 9; c++ {
//            val := board[r][c]
//            if val == '.' {
//                continue
//            }
//
//            bit := 1 << (val - '1')  // '1'→bit 1, '2'→bit 2, etc.
//            boxIdx := (r/3)*3 + (c/3)
//
//            // Check if bit already set
//            if rows[r]&bit != 0 || cols[c]&bit != 0 || boxes[boxIdx]&bit != 0 {
//                return false
//            }
//
//            // Set the bit
//            rows[r] |= bit
//            cols[c] |= bit
//            boxes[boxIdx] |= bit
//        }
//    }
//    return true
//}

func main() {
	// Test case 1: Valid Sudoku
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Printf("Test 1 (Valid Sudoku): %v\n", isValidSudoku(board1))

	// Test case 2: Invalid - duplicate in row
	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Printf("Test 2 (Invalid - duplicate in row): %v\n", isValidSudoku(board2))

	// Test case 3: Invalid - duplicate in column
	board3 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Printf("Test 3 (Invalid - duplicate in column): %v\n", isValidSudoku(board3))

	// Test case 4: Invalid - duplicate in 3x3 box
	board4 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '5', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Printf("Test 4 (Invalid - duplicate in 3x3 box): %v\n", isValidSudoku(board4))
}
