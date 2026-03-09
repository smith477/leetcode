package main

import "fmt"

// func climbStairs(n int) int {

// 	curr, prev1, prev2 := 0, 0, 1

// 	for i := 0; i < n; i++ {
// 		curr = prev1 + prev2
// 		prev1 = prev2
// 		prev2 = curr
// 	}

// 	return curr
// }

//func climbStairs(n int) int {
//
//	prev1, prev2 := 0, 1
//
//	for ; n > 0; n-- {
//		prev1, prev2 = prev2, prev1+prev2
//	}
//
//	return prev2
//}

func climbStairs(n int) int {
	prev, next := 0, 1

	for i := 1; i <= n; i++ {
		prev, next = next, prev+next
	}

	return prev
}

func main() {
	testCases := []struct {
		n        int
		expected int
	}{
		{n: 1, expected: 1},
		{n: 2, expected: 2},
		{n: 3, expected: 3},
		{n: 4, expected: 5},
		{n: 5, expected: 8},
		{n: 6, expected: 13},
		{n: 10, expected: 89},
		{n: 15, expected: 987},
		{n: 20, expected: 10946},
	}

	fmt.Println("Climbing Stairs Test Cases:")
	fmt.Println("===========================")
	for _, tc := range testCases {
		result := climbStairs(tc.n)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("%s n=%d: got %d, expected %d\n", status, tc.n, result, tc.expected)
	}
}
