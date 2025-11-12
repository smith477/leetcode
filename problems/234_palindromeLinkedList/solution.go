package main

import (
	"fmt"
	"leetcode/util"
)

func isPalindrome(head *util.ListNode) bool {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *util.ListNode = nil
	curr := slow

	for curr != nil {
		//Save previous
		next := curr.Next

		// Reverse pointer
		curr.Next = prev

		// Move to the next
		prev = curr
		curr = next
	}

	for prev != nil && head != nil {
		if prev.Val != head.Val {
			return false
		}

		prev, head = prev.Next, head.Next
	}

	return true
}

func main() {
	// Example 1: [1,2,2,1] -> true
	node1 := &util.ListNode{Val: 1}
	node2 := &util.ListNode{Val: 2}
	node3 := &util.ListNode{Val: 2}
	node4 := &util.ListNode{Val: 1}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	fmt.Printf("Example 1 [1,2,2,1]: %v\n", isPalindrome(node1)) // Expected: true

	// Example 2: [1,2] -> false
	node5 := &util.ListNode{Val: 1}
	node6 := &util.ListNode{Val: 2}
	node5.Next = node6

	fmt.Printf("Example 2 [1,2]: %v\n", isPalindrome(node5)) // Expected: false

	// Additional test: [1] -> true
	node7 := &util.ListNode{Val: 1}
	fmt.Printf("Example 3 [1]: %v\n", isPalindrome(node7)) // Expected: true

	// Additional test: [1,2,3,2,1] -> true
	n1 := &util.ListNode{Val: 1}
	n2 := &util.ListNode{Val: 2}
	n3 := &util.ListNode{Val: 3}
	n4 := &util.ListNode{Val: 2}
	n5 := &util.ListNode{Val: 1}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	fmt.Printf("Example 4 [1,2,3,2,1]: %v\n", isPalindrome(n1)) // Expected: true
}
