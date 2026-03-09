package main

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2

	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
