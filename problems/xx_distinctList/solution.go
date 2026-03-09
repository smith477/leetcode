package main

import "fmt"

func distinctList(arr []int) int {
	var count int
	seen := make(map[int]bool)

	for _, el := range arr {
		if seen[el] {
			count++
		} else {
			seen[el] = true
		}
	}

	return count
}

//[1,2,2,2,3]

func main() {
	fmt.Println(distinctList([]int{1, 2, 2, 2, 3})) // 2
	fmt.Println(distinctList([]int{1, 2, 3, 4, 5})) // 0
	fmt.Println(distinctList([]int{1, 1, 1, 1}))    // 3
	fmt.Println(distinctList([]int{}))              // 0
	fmt.Println(distinctList([]int{5, 5, 5, 2, 2})) // 3
}
