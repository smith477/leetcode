package main

// Two-pointer approach on sorted array
// Time: O(n), Space: O(1)
// Use two pointers from both ends, adjust based on sum comparison with target
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			break
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{left + 1, right + 1}
}

func main() {
	// Test case 1
	numbers1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := twoSum(numbers1, target1)
	println("Test 1:", result1[0], result1[1], "Expected: [1, 2]")

	// Test case 2
	numbers2 := []int{2, 3, 4}
	target2 := 6
	result2 := twoSum(numbers2, target2)
	println("Test 2:", result2[0], result2[1], "Expected: [1, 3]")

	// Test case 3
	numbers3 := []int{-1, 0}
	target3 := -1
	result3 := twoSum(numbers3, target3)
	println("Test 3:", result3[0], result3[1], "Expected: [1, 2]")
}
