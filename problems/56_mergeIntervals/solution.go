package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 0; i < len(intervals); i++ {
		last := result[len(result)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			last[1] = max(last[1], current[1])
		} else {
			result = append(result, current)
		}
	}

	return result
}
