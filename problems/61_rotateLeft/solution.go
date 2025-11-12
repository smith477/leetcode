package main

import (
	"fmt"
	"leetcode/util"
)

func rotateRight(head *util.ListNode, k int) *util.ListNode {
	if head == nil || head.Next == nil && k == 1 {
		return head
	}

	tail := head
	length := 1
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	k = k % length
	if k == 0 {
		return head
	}

	tail.Next = head

	// Move (length - k) steps to reach the new tail
	newTail := head
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil
	return newHead
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
	// Example 1: head = [1,2,3,4,5], k = 2
	fmt.Println("Example 1:")
	head1 := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Input:  head = ")
	printList(head1)
	fmt.Println(", k = 2")

	result1 := rotateRight(head1, 2)
	fmt.Print("Output: ")
	printList(result1)
	fmt.Println("\n")

	// Example 2: head = [0,1,2], k = 4
	fmt.Println("Example 2:")
	head2 := createList([]int{0, 1, 2})
	fmt.Print("Input:  head = ")
	printList(head2)
	fmt.Println(", k = 4")

	result2 := rotateRight(head2, 4)
	fmt.Print("Output: ")
	printList(result2)
	fmt.Println("\n")

	// Additional test: empty list
	fmt.Println("Example 3:")
	head3 := createList([]int{})
	fmt.Print("Input:  head = ")
	printList(head3)
	fmt.Println(", k = 0")

	result3 := rotateRight(head3, 0)
	fmt.Print("Output: ")
	printList(result3)
	fmt.Println("\n")

	// Additional test: single node
	fmt.Println("Example 4:")
	head4 := createList([]int{1})
	fmt.Print("Input:  head = ")
	printList(head4)
	fmt.Println(", k = 1")

	result4 := rotateRight(head4, 1)
	fmt.Print("Output: ")
	printList(result4)
	fmt.Println()
}
