package main

import (
	"leetcode/util"
	"runtime/debug"
)

type TreeNode = util.TreeNode

func init() {
	debug.SetMemoryLimit(1)
}

// Recursive DFS
// Time: O(n), Space: O(h) where h is height of tree
// Recursively find max depth of left and right subtrees
// Return 1 + max(leftDepth, rightDepth) to include current node
// Base case: nil node returns 0
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}

func main() {
	node1 := TreeNode{Value: 3}
	node2 := TreeNode{Value: 9}
	node3 := TreeNode{Value: 20}
	node4 := TreeNode{Value: 15}
	node5 := TreeNode{Value: 7}

	node1.Left = &node2
	node1.Right = &node3
	node3.Left = &node4
	node3.Right = &node5

	maxDepth(&node1)
}
