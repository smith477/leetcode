package main

import (
	"fmt"
	"leetcode/util"
)

type Node util.ListNode

func hasCycle(head *util.ListNode) *util.ListNode {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			slow = head
			break
		}
	}

	for slow != nil && fast != nil {
		if slow == fast {
			return slow
		}

		slow = slow.Next
		fast = fast.Next
	}

	return nil
}

func main() {
	// Example 1: [3,2,0,-4], pos = 1
	node1 := &util.ListNode{Val: 3}
	node2 := &util.ListNode{Val: 2}
	node3 := &util.ListNode{Val: 0}
	node4 := &util.ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // Creates cycle at position 1

	fmt.Println("Example 1:")
	fmt.Printf("Output: %v\n", hasCycle(node1))
	fmt.Println("Expected: true\n")

	// Example 2: [1,2], pos = 0
	node5 := &util.ListNode{Val: 1}
	node6 := &util.ListNode{Val: 2}
	node5.Next = node6
	node6.Next = node5 // Creates cycle at position 0

	fmt.Println("Example 2:")
	fmt.Printf("Output: %v\n", hasCycle(node5))
	fmt.Println("Expected: true\n")

	// Example 3: [1], pos = -1
	node7 := &util.ListNode{Val: 1}

	fmt.Println("Example 3:")
	fmt.Printf("Output: %v\n", hasCycle(node7))
	fmt.Println("Expected: false\n")

	node8 := &util.ListNode{Val: 1}
	node9 := &util.ListNode{Val: 2}
	node8.Next = node9

	fmt.Println("Example 4:")
	fmt.Printf("Output: %v\n", hasCycle(node8))
	fmt.Println("Expected: false\n")
}
