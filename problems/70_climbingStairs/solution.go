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

func climbStairs(n int) int {

	prev1, prev2 := 0, 1

	for ; n > 0; n-- {
		prev1, prev2 = prev2, prev1+prev2
	}

	return prev2
}

func main() {

	n := 6
	fmt.Printf("You have %d combination to climb %d stairs\n", climbStairs(n), n)
}
