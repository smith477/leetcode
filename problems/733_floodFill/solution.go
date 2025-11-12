package main

import (
	"fmt"
	"reflect"
)

// func floodFill(image [][]int, sr int, sc int, color int) [][]int {
// 	startColor := image[sr][sc]

// 	if startColor == color {
// 		return image
// 	}

// 	rows := len(image)
// 	cols := len(image[0])

// 	queue := [][]int{{sr, sc}}
// 	image[sr][sc] = color

// 	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 	for len(queue) > 0 {
// 		pixel := queue[0]
// 		queue = queue[1:]

// 		row, col := pixel[0], pixel[1]

// 		for _, dir := range directions {
// 			newRow := row + dir[0]
// 			newCol := col + dir[1]

// 			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && image[newRow][newCol] == startColor {
// 				image[newRow][newCol] = color
// 				queue = append(queue, []int{newRow, newCol})
// 			}
// 		}
// 	}

// 	return image
// }

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	startColor := image[sr][sc]

	if startColor == color {
		return image
	}

	dfs(image, sr, sc, startColor, color)
	return image
}

func dfs(image [][]int, row, col, startColor, newColor int) {

	if row < 0 || row >= len(image) ||
		col < 0 || col >= len(image[0]) ||
		image[row][col] != startColor {
		return
	}

	image[row][col] = newColor

	dfs(image, row-1, col, startColor, newColor)
	dfs(image, row+1, col, startColor, newColor)
	dfs(image, row, col-1, startColor, newColor)
	dfs(image, row, col+1, startColor, newColor)
}

// Helper to copy 2D slice
func copyImage(image [][]int) [][]int {
	result := make([][]int, len(image))
	for i := range image {
		result[i] = make([]int, len(image[i]))
		copy(result[i], image[i])
	}
	return result
}

func main() {
	// Example 1
	fmt.Println("=== Example 1 ===")
	image1 := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	input1 := copyImage(image1)
	result1 := floodFill(image1, 1, 1, 2)
	expected1 := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}

	fmt.Println("Input:   ", input1)
	fmt.Println("Output:  ", result1)
	fmt.Println("Expected:", expected1)
	fmt.Println("Match:", reflect.DeepEqual(result1, expected1))

	// Example 2
	fmt.Println("\n=== Example 2 ===")
	image2 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}
	input2 := copyImage(image2)
	result2 := floodFill(image2, 0, 0, 0)
	expected2 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}

	fmt.Println("Input:   ", input2)
	fmt.Println("Output:  ", result2)
	fmt.Println("Expected:", expected2)
	fmt.Println("Match:", reflect.DeepEqual(result2, expected2))

	// Test 3: Single pixel
	fmt.Println("\n=== Test 3: Single Pixel ===")
	image3 := [][]int{{1}}
	result3 := floodFill(image3, 0, 0, 2)
	expected3 := [][]int{{2}}
	fmt.Println("Output:  ", result3)
	fmt.Println("Expected:", expected3)
	fmt.Println("Match:", reflect.DeepEqual(result3, expected3))

	// Test 4: Separate regions
	fmt.Println("\n=== Test 4: Separate Regions ===")
	image4 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	result4 := floodFill(image4, 1, 1, 2)
	expected4 := [][]int{
		{1, 1, 1},
		{1, 2, 1},
		{1, 1, 1},
	}
	fmt.Println("Output:  ", result4)
	fmt.Println("Expected:", expected4)
	fmt.Println("Match:", reflect.DeepEqual(result4, expected4))

	// Test 5: Large connected region
	fmt.Println("\n=== Test 5: All Connected ===")
	image5 := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	result5 := floodFill(image5, 0, 0, 3)
	expected5 := [][]int{
		{3, 3, 3},
		{3, 3, 3},
		{3, 3, 3},
	}
	fmt.Println("Output:  ", result5)
	fmt.Println("Expected:", expected5)
	fmt.Println("Match:", reflect.DeepEqual(result5, expected5))
}
