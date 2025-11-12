package main

import (
	"fmt"
	"leetcode/util"
)

// Linked Lists
// Time complexity O(n), Space complexity O(1)
// Initialization of prev pointer, the dummy pointer that is pointing to the head of the linked list
// Cycle through the end in a while
// In each iteration initialize first and second as prev.Next, prev.Next.Next (in first iteration that is first and second element of linked list)
// Then make a swap, prev.Next -> second, first.Next -> second.Next, second.Next -> first and finally dummy should be moved to new one dummy = first
// headerPointer.next is the header that should outputted
func swapPairs(head *util.ListNode) *util.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	headPointer := &util.ListNode{Next: head}

	prev := headPointer

	for prev.Next != nil && prev.Next.Next != nil {

		first := prev.Next
		second := prev.Next.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
	}

	return headPointer.Next
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
	// Example 1: [1,2,3,4]
	fmt.Println("Example 1:")
	head1 := createList([]int{1, 2, 3, 4})
	fmt.Print("Input:  ")
	printList(head1)
	fmt.Println()

	result1 := swapPairs(head1)
	fmt.Print("Output: ")
	printList(result1)
	fmt.Println("\n")

	// Example 2: []
	fmt.Println("Example 2:")
	head2 := createList([]int{})
	fmt.Print("Input:  ")
	printList(head2)
	fmt.Println()

	result2 := swapPairs(head2)
	fmt.Print("Output: ")
	printList(result2)
	fmt.Println("\n")

	// Example 3: [1]
	fmt.Println("Example 3:")
	head3 := createList([]int{1})
	fmt.Print("Input:  ")
	printList(head3)
	fmt.Println()

	result3 := swapPairs(head3)
	fmt.Print("Output: ")
	printList(result3)
	fmt.Println("\n")

	// Example 4: [1,2,3]
	fmt.Println("Example 4:")
	head4 := createList([]int{1, 2, 3})
	fmt.Print("Input:  ")
	printList(head4)
	fmt.Println()

	result4 := swapPairs(head4)
	fmt.Print("Output: ")
	printList(result4)
	fmt.Println()
}
