package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			currentLevel = append(currentLevel, node.Value)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
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

	result1 := levelOrder(root1)
	fmt.Println("Output:", result1)
	fmt.Println("Expected: [[3],[9,20],[15,7]]")

	// Example 2: [1]
	fmt.Println("\n=== Example 2 ===")
	root2 := &TreeNode{Value: 1}
	result2 := levelOrder(root2)
	fmt.Println("Output:", result2)
	fmt.Println("Expected: [[1]]")

	// Example 3: []
	fmt.Println("\n=== Example 3 ===")
	result3 := levelOrder(nil)
	fmt.Println("Output:", result3)
	fmt.Println("Expected: []")

	// Additional test: Complete binary tree
	//       1
	//      / \
	//     2   3
	//    / \ / \
	//   4  5 6  7
	fmt.Println("\n=== Additional Test: Complete Tree ===")
	root4 := &TreeNode{Value: 1}
	root4.Left = &TreeNode{Value: 2}
	root4.Right = &TreeNode{Value: 3}
	root4.Left.Left = &TreeNode{Value: 4}
	root4.Left.Right = &TreeNode{Value: 5}
	root4.Right.Left = &TreeNode{Value: 6}
	root4.Right.Right = &TreeNode{Value: 7}

	result4 := levelOrder(root4)
	fmt.Println("Output:", result4)
	fmt.Println("Expected: [[1],[2,3],[4,5,6,7]]")
}
