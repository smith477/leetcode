package main

import "fmt"

// Hash map approach
// Time: O(n), Space: O(n)
// Track count of each number in hash map
// Return the number when its count exceeds n/2
// func majorityElement(nums []int) int {

// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	n := len(nums)

// 	numbersMap := make(map[int]int)

// 	for _, number := range nums {

// 		if count, found := numbersMap[number]; found {
// 			if count >= n/2 {
// 				return number
// 			}
// 		}

// 		numbersMap[number] += 1
// 	}

// 	return -1
// }

// Boyer-Moore Voting Algorithm
// Time: O(n), Space: O(1)
// Track candidate majority element and its count
// Increment count when seeing candidate, decrement when seeing different element
// When count reaches 0, switch candidate to current element
// Guaranteed to find majority element (appears > n/2 times)
func majorityElement(nums []int) int {
	majorElement, count := 0, 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			count++
			majorElement = nums[i]
		} else if majorElement == nums[i] {
			count++
		} else {
			count--
		}
	}

	return majorElement
}

func main() {
	// Test Case 1
	nums1 := []int{3, 2, 3}
	fmt.Printf("Input: %v\n", nums1)
	fmt.Printf("Output: %v\n", majorityElement(nums1))
	fmt.Printf("Expected: 3\n\n")

	// Test Case 2
	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("Input: %v\n", nums2)
	fmt.Printf("Output: %v\n", majorityElement(nums2))
	fmt.Printf("Expected: 2\n\n")

	// Additional Test Cases
	nums3 := []int{1}
	fmt.Printf("Input: %v\n", nums3)
	fmt.Printf("Output: %v\n", majorityElement(nums3))
	fmt.Printf("Expected: 1\n\n")

	nums4 := []int{6, 5, 5}
	fmt.Printf("Input: %v\n", nums4)
	fmt.Printf("Output: %v\n", majorityElement(nums4))
	fmt.Printf("Expected: 5\n\n")
}
