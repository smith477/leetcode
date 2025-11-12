package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// Helper function to print tree in level order (BFS)
func printLevelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func main() {
	// Example 1: [4,2,7,1,3,6,9]
	//       4
	//      / \
	//     2   7
	//    / \ / \
	//   1  3 6  9
	fmt.Println("=== Example 1 ===")
	root1 := &TreeNode{Value: 4}
	root1.Left = &TreeNode{Value: 2}
	root1.Right = &TreeNode{Value: 7}
	root1.Left.Left = &TreeNode{Value: 1}
	root1.Left.Right = &TreeNode{Value: 3}
	root1.Right.Left = &TreeNode{Value: 6}
	root1.Right.Right = &TreeNode{Value: 9}

	fmt.Println("Before:", printLevelOrder(root1))
	invertTree(root1)
	fmt.Println("After: ", printLevelOrder(root1))
	fmt.Println("Expected: [4,7,2,9,6,3,1]")

	// Example 2: [2,1,3]
	//     2
	//    / \
	//   1   3
	fmt.Println("\n=== Example 2 ===")
	root2 := &TreeNode{Value: 2}
	root2.Left = &TreeNode{Value: 1}
	root2.Right = &TreeNode{Value: 3}

	fmt.Println("Before:", printLevelOrder(root2))
	invertTree(root2)
	fmt.Println("After: ", printLevelOrder(root2))
	fmt.Println("Expected: [2,3,1]")

	// Example 3: []
	fmt.Println("\n=== Example 3 ===")
	root3 := invertTree(nil)
	fmt.Println("Before: []")
	fmt.Println("After: ", printLevelOrder(root3))
	fmt.Println("Expected: []")

	// Additional test: Single node
	fmt.Println("\n=== Additional Test: Single Node ===")
	root4 := &TreeNode{Value: 1}
	fmt.Println("Before:", printLevelOrder(root4))
	invertTree(root4)
	fmt.Println("After: ", printLevelOrder(root4))

	// Additional test: Left-skewed tree
	//   1
	//  /
	// 2
	//  /
	// 3
	fmt.Println("\n=== Additional Test: Left-Skewed ===")
	root5 := &TreeNode{Value: 1}
	root5.Left = &TreeNode{Value: 2}
	root5.Left.Left = &TreeNode{Value: 3}

	fmt.Println("Before:", printLevelOrder(root5))
	invertTree(root5)
	fmt.Println("After: ", printLevelOrder(root5))
}
