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

	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			// Skip duplicate
			curr.Next = curr.Next.Next
		} else {
			// Move forward
			curr = curr.Next
		}
	}
	return head
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
