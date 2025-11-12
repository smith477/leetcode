package main

import "fmt"

func isHappy(n int) bool {

	numberMap := make(map[int]bool)
	for n != 1 {
		if numberMap[n] {
			return false
		}

		numberMap[n] = true
		localResult := 0
		for n > 0 {
			reminder := n % 10
			localResult += reminder * reminder
			n /= 10
		}

		n = localResult
		fmt.Printf("Currnet n is %d\n", n)
	}
	return true
}

func main() {
	fmt.Printf("Currnet n is %v\n", isHappy(19))
	fmt.Printf("Currnet n is %v\n", isHappy(2))
}
