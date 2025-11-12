package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

// Middle -> Left -> Right
func dfsPreorder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.Value, " ")
	dfsPreorder(root.Left)
	dfsPreorder(root.Right)
}

// Left -> Middle -> Right
func dfsInorder(root *TreeNode) {
	if root == nil {
		return
	}

	dfsPreorder(root.Left)
	fmt.Print(root.Value, " ")
	dfsPreorder(root.Right)
}

// Middle -> Left -> Right
func dfsPostorder(root *TreeNode) {
	if root == nil {
		return
	}

	dfsPreorder(root.Left)
	dfsPreorder(root.Right)
	fmt.Print(root.Value, " ")
}

// Iterative - Using Stack
func dfsIterative(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Print(node.Value, " ")

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

// DFS with collection - returns slice
func dfsCollect(root *TreeNode) []int {
	result := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Value)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

// BFS - Using Queue
func bfsIterative(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node.Value, " ")

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	// Test Case 1: Regular tree
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5

	fmt.Println("=== Test Case 1: Regular Tree ===")
	root1 := &TreeNode{Value: 1}
	root1.Left = &TreeNode{Value: 2}
	root1.Right = &TreeNode{Value: 3}
	root1.Left.Left = &TreeNode{Value: 4}
	root1.Left.Right = &TreeNode{Value: 5}

	fmt.Print("Preorder:  ")
	dfsPreorder(root1)
	fmt.Println()

	fmt.Print("Inorder:   ")
	dfsInorder(root1)
	fmt.Println()

	fmt.Print("Postorder: ")
	dfsPostorder(root1)
	fmt.Println()

	fmt.Print("Iterative: ")
	dfsIterative(root1)
	fmt.Println()

	fmt.Print("BFS: ")
	bfsIterative(root1)
	fmt.Println()

	fmt.Println("Collected:", dfsCollect(root1))

	// Test Case 2: Single node
	fmt.Println("\n=== Test Case 2: Single Node ===")
	root2 := &TreeNode{Value: 42}
	fmt.Print("Preorder: ")
	dfsPreorder(root2)
	fmt.Println()

	// Test Case 3: Left-skewed tree
	//   1
	//  /
	// 2
	//  /
	// 3
	fmt.Println("\n=== Test Case 3: Left-Skewed Tree ===")
	root3 := &TreeNode{Value: 1}
	root3.Left = &TreeNode{Value: 2}
	root3.Left.Left = &TreeNode{Value: 3}

	fmt.Print("Preorder:  ")
	dfsPreorder(root3)
	fmt.Println()

	// Test Case 4: Right-skewed tree
	// 1
	//  \
	//   2
	//    \
	//     3
	fmt.Println("\n=== Test Case 4: Right-Skewed Tree ===")
	root4 := &TreeNode{Value: 1}
	root4.Right = &TreeNode{Value: 2}
	root4.Right.Right = &TreeNode{Value: 3}

	fmt.Print("Preorder:  ")
	dfsPreorder(root4)
	fmt.Println()

	// Test Case 5: Empty tree
	fmt.Println("\n=== Test Case 5: Empty Tree ===")
	fmt.Print("Preorder: ")
	dfsPreorder(nil)
	fmt.Println("(empty)")

	// Test Case 6: Complete binary tree
	//       1
	//      / \
	//     2   3
	//    / \ / \
	//   4  5 6  7
	fmt.Println("\n=== Test Case 6: Complete Binary Tree ===")
	root6 := &TreeNode{Value: 1}
	root6.Left = &TreeNode{Value: 2}
	root6.Right = &TreeNode{Value: 3}
	root6.Left.Left = &TreeNode{Value: 4}
	root6.Left.Right = &TreeNode{Value: 5}
	root6.Right.Left = &TreeNode{Value: 6}
	root6.Right.Right = &TreeNode{Value: 7}

	fmt.Print("Preorder:  ")
	dfsPreorder(root6)
	fmt.Println()

	fmt.Print("Inorder:   ")
	dfsInorder(root6)
	fmt.Println()

	fmt.Print("Postorder: ")
	dfsPostorder(root6)
	fmt.Println()

	fmt.Print("Iterative: ")
	dfsIterative(root6)
	fmt.Println()
}
