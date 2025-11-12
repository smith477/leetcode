package main

import (
	"fmt"
	"leetcode/util"
)

func reverseList(head *util.ListNode) *util.ListNode {

	var prev *util.ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}

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
	// Example 1: [1,2,3,4,5]
	fmt.Println("Example 1:")
	head1 := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Input:  ")
	printList(head1)
	fmt.Println()

	reversed1 := reverseList(head1)
	fmt.Print("Output: ")
	printList(reversed1)
	fmt.Println("\n")

	// Example 2: [1,2]
	fmt.Println("Example 2:")
	head2 := createList([]int{1, 2})
	fmt.Print("Input:  ")
	printList(head2)
	fmt.Println()

	reversed2 := reverseList(head2)
	fmt.Print("Output: ")
	printList(reversed2)
	fmt.Println("\n")

	// Example 3: []
	fmt.Println("Example 3:")
	head3 := createList([]int{})
	fmt.Print("Input:  ")
	printList(head3)
	fmt.Println()

	reversed3 := reverseList(head3)
	fmt.Print("Output: ")
	printList(reversed3)
	fmt.Println()
}
