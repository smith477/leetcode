package main

import (
	"fmt"
	"leetcode/util"
	"math"
)

type TreeNode = util.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := []float64{}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levelSum += node.Value

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelAverage := float64(levelSum) / float64(levelSize)
		result = append(result, levelAverage)
	}

	return result
}

// Helper to compare float slices with tolerance
func floatSlicesEqual(a, b []float64, tolerance float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > tolerance {
			return false
		}
	}
	return true
}

func main() {
	// Example 1: [3,9,20,null,null,15,7]
	//       3
	//      / \
	//     9  20
	//       /  \
	//      15   7
	fmt.Println("=== Example 1 ===")
	root1 := &TreeNode{Value: 3}
	root1.Left = &TreeNode{Value: 9}
	root1.Right = &TreeNode{Value: 20}
	root1.Right.Left = &TreeNode{Value: 15}
	root1.Right.Right = &TreeNode{Value: 7}

	result1 := averageOfLevels(root1)
	expected1 := []float64{3.00000, 14.50000, 11.00000}
	fmt.Printf("Output:   %v\n", result1)
	fmt.Printf("Expected: %v\n", expected1)
	fmt.Printf("Match: %v\n", floatSlicesEqual(result1, expected1, 0.00001))

	// Example 2: [3,9,20,15,7]
	//       3
	//      / \
	//     9  20
	//    / \
	//   15  7
	fmt.Println("\n=== Example 2 ===")
	root2 := &TreeNode{Value: 3}
	root2.Left = &TreeNode{Value: 9}
	root2.Right = &TreeNode{Value: 20}
	root2.Left.Left = &TreeNode{Value: 15}
	root2.Left.Right = &TreeNode{Value: 7}

	result2 := averageOfLevels(root2)
	expected2 := []float64{3.00000, 14.50000, 11.00000}
	fmt.Printf("Output:   %v\n", result2)
	fmt.Printf("Expected: %v\n", expected2)
	fmt.Printf("Match: %v\n", floatSlicesEqual(result2, expected2, 0.00001))

	// Test 3: Single node
	fmt.Println("\n=== Test 3: Single Node ===")
	root3 := &TreeNode{Value: 5}
	result3 := averageOfLevels(root3)
	expected3 := []float64{5.00000}
	fmt.Printf("Output:   %v\n", result3)
	fmt.Printf("Expected: %v\n", expected3)
	fmt.Printf("Match: %v\n", floatSlicesEqual(result3, expected3, 0.00001))

	// Test 4: Complete binary tree
	//       1
	//      / \
	//     2   3
	//    / \ / \
	//   4  5 6  7
	fmt.Println("\n=== Test 4: Complete Binary Tree ===")
	root4 := &TreeNode{Value: 1}
	root4.Left = &TreeNode{Value: 2}
	root4.Right = &TreeNode{Value: 3}
	root4.Left.Left = &TreeNode{Value: 4}
	root4.Left.Right = &TreeNode{Value: 5}
	root4.Right.Left = &TreeNode{Value: 6}
	root4.Right.Right = &TreeNode{Value: 7}

	result4 := averageOfLevels(root4)
	expected4 := []float64{1.00000, 2.50000, 5.50000}
	fmt.Printf("Output:   %v\n", result4)
	fmt.Printf("Expected: %v\n", expected4)
	fmt.Printf("Match: %v\n", floatSlicesEqual(result4, expected4, 0.00001))

	// Test 5: Left-skewed tree
	//   1
	//  /
	// 2
	//  /
	// 3
	fmt.Println("\n=== Test 5: Left-Skewed Tree ===")
	root5 := &TreeNode{Value: 1}
	root5.Left = &TreeNode{Value: 2}
	root5.Left.Left = &TreeNode{Value: 3}

	result5 := averageOfLevels(root5)
	expected5 := []float64{1.00000, 2.00000, 3.00000}
	fmt.Printf("Output:   %v\n", result5)
	fmt.Printf("Expected: %v\n", expected5)
	fmt.Printf("Match: %v\n", floatSlicesEqual(result5, expected5, 0.00001))

	// Test 6: Tree with negative values
	//      -1
	//      / \
	//    -2  -3
	fmt.Println("\n=== Test 6: Negative Values ===")
	root6 := &TreeNode{Value: -1}
	root6.Left = &TreeNode{Value: -2}
	root6.Right = &TreeNode{Value: -3}

	result6 := averageOfLevels(root6)
	expected6 := []float64{-1.00000, -2.50000}
	fmt.Printf("Output:   %v\n", result6)
	fmt.Printf("Expected: %v\n", expected6)
	fmt.Printf("Match: %v\n", floatSlicesEqual(result6, expected6, 0.00001))

	// Test 7: Large values (Int overflow test)
	//      2147483647
	//      /        \
	// 2147483647  2147483647
	fmt.Println("\n=== Test 7: Large Values (Int Overflow Test) ===")
	root7 := &TreeNode{Value: 2147483647}
	root7.Left = &TreeNode{Value: 2147483647}
	root7.Right = &TreeNode{Value: 2147483647}

	result7 := averageOfLevels(root7)
	expected7 := []float64{2147483647.00000, 2147483647.00000}
	fmt.Printf("Output:   %v\n", result7)
	fmt.Printf("Expected: %v\n", expected7)
	fmt.Printf("Match: %v\n", floatSlicesEqual(result7, expected7, 0.00001))
}
