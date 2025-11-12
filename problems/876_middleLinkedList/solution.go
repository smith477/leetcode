package main

import (
	"fmt"
	"leetcode/util"
)

// func middleNode(head *util.ListNode) *util.ListNode {

// 	start := head

// 	count := 0
// 	for start != nil {
// 		start = start.Next
// 		count++
// 	}

// 	start = head
// 	for index := 0; index < count/2; index++ {
// 		start = start.Next
// 	}

// 	return start
// }

func middleNode(head *util.ListNode) *util.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	// Example 1: [1,2,3,4,5] -> should return node 3
	node1 := &util.ListNode{Val: 1}
	node2 := &util.ListNode{Val: 2}
	node3 := &util.ListNode{Val: 3}
	node4 := &util.ListNode{Val: 4}
	node5 := &util.ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	result := middleNode(node1)
	fmt.Printf("Example 1: %d\n", result.Val) // Expected: 3

	// Example 2: [1,2,3,4,5,6] -> should return node 4
	node6 := &util.ListNode{Val: 6}
	node5.Next = node6

	result = middleNode(node1)
	fmt.Printf("Example 2: %d\n", result.Val) // Expected: 4
}
