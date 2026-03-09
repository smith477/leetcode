package main

import "math"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n))) + 1
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
