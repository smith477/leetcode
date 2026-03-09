package main

import "leetcode/util"

func removeNthFromEnd(head *util.ListNode, n int) *util.ListNode {
	dummyNode := &util.ListNode{Next: head}

	leaderNode := dummyNode
	followerNode := dummyNode

	// Move leader n+1 steps ahead to create proper gap
	gapCounter := 0
	for gapCounter <= n {
		leaderNode = leaderNode.Next
		gapCounter++
	}

	// Move both pointers until leader reaches end
	for leaderNode != nil {
		leaderNode = leaderNode.Next
		followerNode = followerNode.Next
	}

	// Remove the nth node from end
	followerNode.Next = followerNode.Next.Next

	return dummyNode.Next
}
