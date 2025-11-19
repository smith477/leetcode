package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
	currentFuel := 0
	startingStation := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currentFuel += gas[i] - cost[i]

		// If we can't reach next station from current start
		if currentFuel < 0 {
			// Skip to next potential starting station
			startingStation = i + 1
			currentFuel = 0 // Reset fuel
		}
	}

	// If total gas < total cost, impossible
	if totalGas < totalCost {
		return -1
	}

	return startingStation
}

func main() {
	testCases := []struct {
		name     string
		gas      []int
		cost     []int
		expected int
	}{
		{
			name:     "Example 1",
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			name:     "Example 2",
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 3},
			expected: -1,
		},
		{
			name:     "Single station - possible",
			gas:      []int{5},
			cost:     []int{4},
			expected: 0,
		},
		{
			name:     "Single station - impossible",
			gas:      []int{2},
			cost:     []int{3},
			expected: -1,
		},
		{
			name:     "All equal gas and cost",
			gas:      []int{3, 3, 3, 3},
			cost:     []int{3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "Two stations - possible",
			gas:      []int{3, 5},
			cost:     []int{4, 3},
			expected: 1,
		},
		{
			name:     "Two stations - impossible",
			gas:      []int{1, 2},
			cost:     []int{2, 2},
			expected: -1,
		},
		{
			name:     "Large surplus at end",
			gas:      []int{1, 1, 1, 10},
			cost:     []int{2, 2, 2, 1},
			expected: 3,
		},
		{
			name:     "Zero gas at some stations",
			gas:      []int{0, 5, 3, 1},
			cost:     []int{1, 2, 3, 2},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		result := canCompleteCircuit(tc.gas, tc.cost)
		passed := result == tc.expected

		status := "✓ PASS"
		if !passed {
			status = "✗ FAIL"
		}

		fmt.Printf("%s - %s\n", status, tc.name)
		if !passed {
			fmt.Printf("  Gas:      %v\n", tc.gas)
			fmt.Printf("  Cost:     %v\n", tc.cost)
			fmt.Printf("  Expected: %d\n", tc.expected)
			fmt.Printf("  Got:      %d\n", result)
		}
	}
}
