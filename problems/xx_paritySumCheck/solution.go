package main

import "fmt"

// ## Problem Statement

// Given an integer array `nums`, you need to:

// 1. **Calculate the difference** between the sum of even-valued elements and odd-valued elements
//    - `evenSum` = sum of all even numbers in the array
//    - `oddSum` = sum of all odd numbers in the array
//    - `diff` = `evenSum - oddSum`

// 2. **Determine divisibility flag** based on `diff` (not total sum):
//    - **"DIV_BY_6"** - if `diff` is divisible by 6 (divisible by both 2 and 3)
//    - **"DIV_BY_3_ONLY"** - if `diff` is divisible by 3 but NOT by 2
//    - **"DIV_BY_2_ONLY"** - if `diff` is divisible by 2 but NOT by 3
//    - **"NONE"** - otherwise

// 3. **Return**: `[diff, flag]`

// ---

// ## Examples

// ### Example 1
// ```
// Input:  nums = [1, 2, 3, 4]
// Output: [2, "DIV_BY_2_ONLY"]

// Explanation:
// - evenSum = 2 + 4 = 6
// - oddSum = 1 + 3 = 4
// - diff = 6 - 4 = 2
// - 2 % 2 == 0 and 2 % 3 != 0 → "DIV_BY_2_ONLY"
// ```

// ### Example 2
// ```
// Input:  nums = [3, 3, 3]
// Output: [-9, "DIV_BY_3_ONLY"]

// Explanation:
// - evenSum = 0
// - oddSum = 3 + 3 + 3 = 9
// - diff = 0 - 9 = -9
// - -9 % 3 == 0 and -9 % 2 != 0 → "DIV_BY_3_ONLY"
// ```

// ### Example 3
// ```
// Input:  nums = [6, 6]
// Output: [12, "DIV_BY_6"]

// Explanation:
// - evenSum = 6 + 6 = 12
// - oddSum = 0
// - diff = 12 - 0 = 12
// - 12 % 6 == 0 → "DIV_BY_6"
// ```

// ---

// ## Constraints

// - `1 <= nums.length <= 10^5`
// - `-10^9 <= nums[i] <= 10^9`

type DivisibilityFlag string

const (
	DivTwo   = "DIV_BY_2_ONLY"
	DivThree = "DIV_BY_3_ONLY"
	DivSix   = "DIV_BY_6_ONLY"
	None     = "None"
)

type parityResult struct {
	Diff int
	Flag DivisibilityFlag
}

func paritySumDifference(nums []int) parityResult {

	oddSum := 0
	evenSum := 0

	for _, element := range nums {

		if element%2 == 0 {
			evenSum += element
		} else {
			oddSum += element
		}
	}

	diff := evenSum - oddSum

	if diff%6 == 0 {
		return parityResult{diff, DivSix}
	} else if diff%3 == 0 {
		return parityResult{diff, DivThree}
	} else if diff%2 == 0 {
		return parityResult{diff, DivTwo}
	} else {
		return parityResult{diff, None}
	}
}

func main() {
	testCases := []struct {
		input    []int
		expected parityResult
	}{
		// Original test cases
		{[]int{1, 2, 3, 4}, parityResult{2, DivTwo}},
		{[]int{3, 3, 3}, parityResult{-9, DivThree}},
		{[]int{6, 6}, parityResult{12, DivSix}},

		// Edge case: None case
		{[]int{1, 2}, parityResult{1, None}}, // diff=1
		{[]int{5}, parityResult{-5, None}},   // diff=-5

		// Edge case: Single element
		{[]int{2}, parityResult{2, DivTwo}},    // only even
		{[]int{3}, parityResult{-3, DivThree}}, // only odd
		{[]int{6}, parityResult{6, DivSix}},    // single element div by 6

		// Edge case: Zero
		{[]int{0}, parityResult{0, DivSix}}, // 0 is even, 0 % 6 == 0
		{[]int{0, 0, 0}, parityResult{0, DivSix}},

		// Edge case: All even
		{[]int{2, 4, 6, 8}, parityResult{20, DivTwo}}, // 20 % 2 == 0, 20 % 3 != 0

		// Edge case: All odd
		{[]int{1, 3, 5, 7}, parityResult{-16, DivTwo}}, // -16 % 2 == 0, -16 % 3 != 0

		// Edge case: Negative numbers
		{[]int{-2, -4}, parityResult{-6, DivSix}}, // negative evens
		{[]int{-1, -3}, parityResult{4, DivTwo}},  // negative odds, diff = 0 - (-4) = 4
		{[]int{-6, 6}, parityResult{0, DivSix}},   // cancels out

		// Edge case: Mixed negative and positive
		{[]int{-2, 2, -3, 3}, parityResult{0, DivSix}}, // everything cancels
		{[]int{10, -5, 3}, parityResult{12, DivSix}},   // 10 - (-5+3) = 12

		// Edge case: Large numbers (within constraints)
		{[]int{1000000000, 5}, parityResult{999999995, None}},
	}

	for index, testCase := range testCases {
		result := paritySumDifference(testCase.input)
		fmt.Printf("Test %d - Expected: %+v, Got: %+v\n", index+1, testCase.expected, result)
	}
}
