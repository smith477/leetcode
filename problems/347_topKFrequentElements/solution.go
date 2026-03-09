package main

func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int)

	for _, num := range nums {
		frequencyMap[num]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, count := range frequencyMap {
		buckets[count] = append(buckets[count], num)
	}

	result := []int{}

	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}
