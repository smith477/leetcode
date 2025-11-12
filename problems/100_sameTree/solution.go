package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	if p.Value != q.Value {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	// Test 1: Same trees
	p1 := &TreeNode{Value: 1}
	p1.Left = &TreeNode{Value: 2}
	p1.Right = &TreeNode{Value: 3}

	q1 := &TreeNode{Value: 1}
	q1.Left = &TreeNode{Value: 2}
	q1.Right = &TreeNode{Value: 3}

	fmt.Println("Test 1:", isSameTree(p1, q1)) // true

	// Test 2: Different structure
	p2 := &TreeNode{Value: 1}
	p2.Left = &TreeNode{Value: 2}

	q2 := &TreeNode{Value: 1}
	q2.Right = &TreeNode{Value: 2}

	fmt.Println("Test 2:", isSameTree(p2, q2)) // false

	// Test 3: Different Valueues
	p3 := &TreeNode{Value: 1}
	p3.Left = &TreeNode{Value: 2}
	p3.Right = &TreeNode{Value: 1}

	q3 := &TreeNode{Value: 1}
	q3.Left = &TreeNode{Value: 1}
	q3.Right = &TreeNode{Value: 2}

	fmt.Println("Test 3:", isSameTree(p3, q3)) // false

	// Test 4: Both empty
	fmt.Println("Test 4:", isSameTree(nil, nil)) // true

	// Test 5: One empty
	p5 := &TreeNode{Value: 1}
	fmt.Println("Test 5:", isSameTree(p5, nil)) // false
}
