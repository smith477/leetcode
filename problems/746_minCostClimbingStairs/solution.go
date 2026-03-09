package main

func minCostClimbingStairs(cost []int) int {
	totalCost := 0

	first, second := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		totalCost += min(first, second)
		first, second = second, min(second, cost[i])
	}

	return totalCost
}
