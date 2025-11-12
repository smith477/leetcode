package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return [][]int{}
	}

	sort.Ints(nums)
	fmt.Println("Sorted nums:", nums)

	result := [][]int{}

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, n-1

			for left < right {
				a, b, c, d := nums[i], nums[j], nums[left], nums[right]
				sum := a + b + c + d

				if sum == target {
					result = append(result, []int{a, b, c, d})

					for left < right && nums[left] == nums[left+1] {
						left++
					}

					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

func main() {
	// Example 1
	nums1 := []int{1, 0, -1, 0, -2, 2}
	target1 := 0
	fmt.Println("Example 1 Output:", fourSum(nums1, target1))

	// Example 2
	nums2 := []int{2, 2, 2, 2, 2}
	target2 := 8
	fmt.Println("Example 2 Output:", fourSum(nums2, target2))
}
