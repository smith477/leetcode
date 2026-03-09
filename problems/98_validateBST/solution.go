package main

import (
	"leetcode/util"
	"math"
)

type TreeNode = util.TreeNode

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt, math.MaxInt)
}

func validate(node *TreeNode, lower int, upper int) bool {
	if node == nil {
		return true
	}

	if lower < node.Value && node.Value < upper {
		return validate(node.Left, lower, node.Value) && validate(node.Right, node.Value, upper)
	} else {
		return false
	}
}
