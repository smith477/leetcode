package main

import "fmt"

// Two pointer pattern
// Time complexity O(n), Space complexity O(1)
// Hash map number:index
// target - mapNumber = number to get the two indexes
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for numIndex, number := range nums {
		if mapIndex, found := numsMap[target-number]; found {
			return []int{mapIndex, numIndex}
		}
		numsMap[number] = numIndex
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
