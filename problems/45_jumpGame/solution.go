package main

import "fmt"

// Greedy approach with implicit BFS levels
// Time: O(n), Space: O(1)
// Track the farthest position reachable in current "jump level"
// When reaching end of current level, increment jumps and move to next level
// Each level represents all positions reachable with same number of jumps
func jump(nums []int) int {

	n := len(nums) - 1
	i, maxReachable, lastJumpedPos, jumps := 0, 0, 0, 0

	for lastJumpedPos < n {
		maxReachable = max(maxReachable, nums[i]+i)
		if i == lastJumpedPos {
			lastJumpedPos = maxReachable
			jumps++
		}
		i++
	}

	return jumps
}

func main() {
	// Test 1
	fmt.Println("=== Test 1 ===")
	nums1 := []int{2, 3, 1, 1, 4}
	result1 := jump(nums1)
	fmt.Printf("Input: %v\n", nums1)
	fmt.Printf("Output: %d jumps\n", result1)
	fmt.Printf("Expected: 2 jumps\n")
	fmt.Printf("Explanation: Jump 1 step from 0→1, then 3 steps to last index\n")
	fmt.Printf("Match: %v\n\n", result1 == 2)

	// Test 2
	fmt.Println("=== Test 2 ===")
	nums2 := []int{2, 3, 0, 1, 4}
	result2 := jump(nums2)
	fmt.Printf("Input: %v\n", nums2)
	fmt.Printf("Output: %d jumps\n", result2)
	fmt.Printf("Expected: 2 jumps\n")
	fmt.Printf("Match: %v\n\n", result2 == 2)

	// Test 3: Single element
	fmt.Println("=== Test 3: Single Element ===")
	nums3 := []int{0}
	result3 := jump(nums3)
	fmt.Printf("Input: %v\n", nums3)
	fmt.Printf("Output: %d jumps\n", result3)
	fmt.Printf("Expected: 0 jumps\n")
	fmt.Printf("Match: %v\n\n", result3 == 0)

	// Test 4: Two elements
	fmt.Println("=== Test 4: Two Elements ===")
	nums4 := []int{1, 1}
	result4 := jump(nums4)
	fmt.Printf("Input: %v\n", nums4)
	fmt.Printf("Output: %d jumps\n", result4)
	fmt.Printf("Expected: 1 jump\n")
	fmt.Printf("Match: %v\n\n", result4 == 1)

	// Test 5: Can jump directly
	fmt.Println("=== Test 5: Direct Jump ===")
	nums5 := []int{5, 1, 1, 1, 1}
	result5 := jump(nums5)
	fmt.Printf("Input: %v\n", nums5)
	fmt.Printf("Output: %d jumps\n", result5)
	fmt.Printf("Expected: 1 jump\n")
	fmt.Printf("Match: %v\n\n", result5 == 1)

	// Test 6: All ones
	fmt.Println("=== Test 6: All Ones ===")
	nums6 := []int{1, 1, 1, 1, 1}
	result6 := jump(nums6)
	fmt.Printf("Input: %v\n", nums6)
	fmt.Printf("Output: %d jumps\n", result6)
	fmt.Printf("Expected: 4 jumps\n")
	fmt.Printf("Match: %v\n\n", result6 == 4)

	// Test 7: Longer array
	fmt.Println("=== Test 7: Longer Array ===")
	nums7 := []int{1, 2, 1, 1, 1}
	result7 := jump(nums7)
	fmt.Printf("Input: %v\n", nums7)
	fmt.Printf("Output: %d jumps\n", result7)
	fmt.Printf("Expected: 3 jumps\n")
	fmt.Printf("Path: 0→1→3→4\n")
	fmt.Printf("Match: %v\n\n", result7 == 3)

	// Test 8: Large jumps
	fmt.Println("=== Test 8: Large Jumps ===")
	nums8 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1}
	result8 := jump(nums8)
	fmt.Printf("Input: %v\n", nums8)
	fmt.Printf("Output: %d jumps\n", result8)
	fmt.Printf("Expected: 1 jump\n")
	fmt.Printf("Match: %v\n\n", result8 == 1)
}
