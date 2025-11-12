package main

import (
	"fmt"
	"leetcode/util"
)

func reverseBetween(head *util.ListNode, left int, right int) *util.ListNode {
	if head == nil || left == right {
		return head
	}

	headPointer := &util.ListNode{Next: head}
	leftPointer := headPointer

	for index := 1; index < left; index++ {
		leftPointer = leftPointer.Next
	}

	var prev *util.ListNode
	curr := leftPointer.Next
	tail := curr

	for index := left; index <= right; index++ {
		next := curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	leftPointer.Next = prev
	tail.Next = curr

	return headPointer.Next
}

// Helper function to create a linked list from a slice
func createList(values []int) *util.ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &util.ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &util.ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

// Helper function to print the linked list
func printList(head *util.ListNode) {
	fmt.Print("[")
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(",")
		}
		current = current.Next
	}
	fmt.Print("]")
}

func main() {
	// Example 1: head = [1,2,3,4,5], left = 2, right = 4
	fmt.Println("Example 1:")
	head1 := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Input:  head = ")
	printList(head1)
	fmt.Println(", left = 2, right = 4")

	result1 := reverseBetween(head1, 2, 4)
	fmt.Print("Output: ")
	printList(result1)
	fmt.Println("\n")

	// Example 2: head = [5], left = 1, right = 1
	fmt.Println("Example 2:")
	head2 := createList([]int{5})
	fmt.Print("Input:  head = ")
	printList(head2)
	fmt.Println(", left = 1, right = 1")

	result2 := reverseBetween(head2, 1, 1)
	fmt.Print("Output: ")
	printList(result2)
	fmt.Println()
}
