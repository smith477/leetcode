package main

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	result := [][]int{}
	i, j := 0, 0

	for i < j {

		start := max(firstList[i][0], secondList[j][0])
		end := min(firstList[i][1], secondList[j][1])

		if start <= end {
			result = append(result, []int{start, end})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return result
}
