package main

import "fmt"

func duplicateZeros(arr []int) {
	zeros := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeros++
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if zeros+i < len(arr) {
				arr[zeros+i] = 0
			}
			if zeros-1+i < len(arr) {
				arr[zeros-1+i] = 0
			}
			zeros--
		} else if i+zeros < len(arr) {
			arr[zeros+i] = arr[i]
		}
	}
}

func main() {
	arr1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	fmt.Println(arr1) // Expected: [1,0,0,2,3,0,0,4]

	arr2 := []int{1, 2, 3}
	duplicateZeros(arr2)
	fmt.Println(arr2) // Expected: [1,2,3]
}
