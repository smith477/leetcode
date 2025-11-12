package main

import "fmt"

var badVersion int

func isBadVersion(version int) bool {
	return version >= badVersion
}

func firstBadVersion(n int) int {

	if isBadVersion(n) {
		return n
	}

	start, end := 1, n

	for start <= end {

		mid := start + (end-start)/2

		if isBadVersion(mid) {
			return mid
		} else if badVersion > mid {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func main() {
	badVersion = 4
	fmt.Printf("n = 5, bad = 4\n")
	fmt.Printf("Output: %d\n\n", firstBadVersion(5))

	badVersion = 1
	fmt.Printf("n = 1, bad = 1\n")
	fmt.Printf("Output: %d\n", firstBadVersion(1))
}
