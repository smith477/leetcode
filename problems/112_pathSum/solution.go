package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

// Recursive DFS
// Time: O(n), Space: O(h) where h is height of tree
// Go deep into tree subtracting current node value from targetSum at each step
// When reaching a leaf node (no left/right children), check if targetSum equals leaf's value
// If match found at any leaf, path exists and is reachable
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Value
	}

	return hasPathSum(root.Left, targetSum-root.Value) ||
		hasPathSum(root.Right, targetSum-root.Value)
}

func hasPathSumIterative(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	sumsMap := map[*TreeNode]int{root: root.Value}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		remainingSum := sumsMap[node]

		if node.Left == nil && node.Right == nil {
			if node.Value == remainingSum {
				return true
			}
			continue
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			sumsMap[node.Left] = targetSum - node.Value
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			sumsMap[node.Right] = targetSum - node.Value
		}
	}

	return false
}

// BFS with Queue
func hasPathSumBFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	type Pair struct {
		node         *TreeNode
		remainingSum int
	}

	queue := []Pair{{root, targetSum}}

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]

		node := pair.node
		remaining := pair.remainingSum

		if node.Left == nil && node.Right == nil && node.Value == remaining {
			return true
		}

		if node.Left != nil {
			queue = append(queue, Pair{node.Left, remaining - node.Value})
		}
		if node.Right != nil {
			queue = append(queue, Pair{node.Right, remaining - node.Value})
		}
	}

	return false
}

func main() {
	// Build test tree:
	//       5
	//      / \
	//     4   8
	//    /   / \
	//   11  13  4
	//  /  \      \
	// 7    2      1

	root := &TreeNode{Value: 5}
	root.Left = &TreeNode{Value: 4}
	root.Right = &TreeNode{Value: 8}
	root.Left.Left = &TreeNode{Value: 11}
	root.Left.Left.Left = &TreeNode{Value: 7}
	root.Left.Left.Right = &TreeNode{Value: 2}
	root.Right.Left = &TreeNode{Value: 13}
	root.Right.Right = &TreeNode{Value: 4}
	root.Right.Right.Right = &TreeNode{Value: 1}

	fmt.Println("=== Testing targetSum = 22 (should be true) ===")
	fmt.Println("Recursive:", hasPathSum(root, 22))
	fmt.Println("Iterative DFS:", hasPathSumIterative(root, 22))
	fmt.Println("BFS:", hasPathSumBFS(root, 22))

	fmt.Println("\n=== Testing targetSum = 5 (should be false) ===")
	fmt.Println("Recursive:", hasPathSum(root, 5))
	fmt.Println("Iterative DFS:", hasPathSumIterative(root, 5))
	fmt.Println("BFS:", hasPathSumBFS(root, 5))

	fmt.Println("\n=== Testing single node ===")
	single := &TreeNode{Value: 1}
	fmt.Println("Recursive (sum=1):", hasPathSum(single, 1))
	fmt.Println("Iterative (sum=1):", hasPathSumIterative(single, 1))

	fmt.Println("\n=== Testing empty tree ===")
	fmt.Println("Recursive:", hasPathSum(nil, 0))
	fmt.Println("Iterative:", hasPathSumIterative(nil, 0))
}
