package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// Create a fake starter node before head
	dummy := &ListNode{Next: head}
	prev := dummy // the "builder" pointer
	curr := head  // the "reader" pointer

	for curr != nil {
		// Check if current value has duplicates
		if curr.Next != nil && curr.Val == curr.Next.Val {
			// Skip all nodes with this value
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}
			// Cut out the whole group
			prev.Next = curr.Next
		} else {
			// No duplicate, keep this node
			prev = prev.Next
		}
		curr = curr.Next
	}

	return dummy.Next
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, v := range nums[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

// Print linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	tests := [][]int{
		{},                          // Empty list
		{1},                         // Single element
		{1, 1, 1, 1},                // All duplicates
		{1, 2, 3, 4, 5},             // No duplicates
		{1, 1, 2, 3},                // Duplicates at start
		{1, 2, 3, 3},                // Duplicates at end
		{1, 1, 2, 3, 3},             // Mixed duplicates
		{1, 1, 2, 2, 3, 4, 4, 4, 5}, // Larger case
	}

	for _, t := range tests {
		head := buildList(t)
		fmt.Print("Input:  ", t, " -> Output: ")
		printList(deleteDuplicates(head))
	}
}
