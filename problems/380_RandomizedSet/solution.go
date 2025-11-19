package main

import (
	"fmt"
	"math/rand"
)

// Hash map + Dynamic array for O(1) operations
// Map stores value -> index mapping for O(1) lookup
// Array stores actual values for O(1) random access
// Remove: swap with last element, then pop (avoids shifting)
type RandomizedSet struct {
	NumberList []int
	NumbersMap map[int]int
}

func Constructor() RandomizedSet {
	var set RandomizedSet
	set.NumberList = []int{}
	set.NumbersMap = make(map[int]int)
	return set
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, found := rs.NumbersMap[val]; found {
		return false
	}

	rs.NumbersMap[val] = len(rs.NumberList)
	rs.NumberList = append(rs.NumberList, val)

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	index, found := rs.NumbersMap[val]
	if !found {
		return false
	}

	// Get last element
	lastIdx := len(rs.NumberList) - 1
	lastVal := rs.NumberList[lastIdx]

	// Move last element to deleted position
	rs.NumberList[index] = lastVal
	rs.NumbersMap[lastVal] = index

	// Remove last element
	rs.NumberList = rs.NumberList[:lastIdx]
	delete(rs.NumbersMap, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(rs.NumberList))
	return rs.NumberList[randomIndex]
}

// Test helper
func testRandomizedSet() {
	fmt.Println("=== Test 1: Example from problem ===")
	rs := Constructor()

	fmt.Printf("insert(1): %v (expected: true)\n", rs.Insert(1))
	fmt.Printf("remove(2): %v (expected: false)\n", rs.Remove(2))
	fmt.Printf("insert(2): %v (expected: true)\n", rs.Insert(2))

	random1 := rs.GetRandom()
	fmt.Printf("getRandom(): %v (expected: 1 or 2)\n", random1)

	fmt.Printf("remove(1): %v (expected: true)\n", rs.Remove(1))
	fmt.Printf("insert(2): %v (expected: false)\n", rs.Insert(2))

	random2 := rs.GetRandom()
	fmt.Printf("getRandom(): %v (expected: 2)\n", random2)

	// Test 2: Insert duplicates
	fmt.Println("\n=== Test 2: Duplicate Inserts ===")
	rs2 := Constructor()
	fmt.Printf("insert(5): %v (expected: true)\n", rs2.Insert(5))
	fmt.Printf("insert(5): %v (expected: false)\n", rs2.Insert(5))
	fmt.Printf("insert(5): %v (expected: false)\n", rs2.Insert(5))

	// Test 3: Remove non-existent
	fmt.Println("\n=== Test 3: Remove Non-existent ===")
	rs3 := Constructor()
	fmt.Printf("remove(10): %v (expected: false)\n", rs3.Remove(10))
	fmt.Printf("insert(10): %v (expected: true)\n", rs3.Insert(10))
	fmt.Printf("remove(10): %v (expected: true)\n", rs3.Remove(10))
	fmt.Printf("remove(10): %v (expected: false)\n", rs3.Remove(10))

	// Test 4: Multiple elements
	fmt.Println("\n=== Test 4: Multiple Elements ===")
	rs4 := Constructor()
	rs4.Insert(1)
	rs4.Insert(2)
	rs4.Insert(3)
	rs4.Insert(4)
	rs4.Insert(5)

	fmt.Println("Inserted: 1, 2, 3, 4, 5")
	fmt.Printf("remove(3): %v (expected: true)\n", rs4.Remove(3))
	fmt.Printf("remove(3): %v (expected: false)\n", rs4.Remove(3))

	// Get random a few times
	fmt.Println("Getting 10 random values (should be from {1,2,4,5}):")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rs4.GetRandom())
	}
	fmt.Println()

	// Test 5: Insert, remove, insert same element
	fmt.Println("\n=== Test 5: Insert-Remove-Insert Same Element ===")
	rs5 := Constructor()
	fmt.Printf("insert(7): %v (expected: true)\n", rs5.Insert(7))
	fmt.Printf("remove(7): %v (expected: true)\n", rs5.Remove(7))
	fmt.Printf("insert(7): %v (expected: true)\n", rs5.Insert(7))
	fmt.Printf("getRandom(): %v (expected: 7)\n", rs5.GetRandom())

	// Test 6: Randomness distribution (basic check)
	fmt.Println("\n=== Test 6: Randomness Check ===")
	rs6 := Constructor()
	rs6.Insert(10)
	rs6.Insert(20)
	rs6.Insert(30)

	count := make(map[int]int)
	trials := 3000
	for i := 0; i < trials; i++ {
		count[rs6.GetRandom()]++
	}

	fmt.Printf("After %d trials:\n", trials)
	for val, cnt := range count {
		percentage := float64(cnt) / float64(trials) * 100
		fmt.Printf("  Value %d: %d times (%.1f%%, expected ~33.3%%)\n", val, cnt, percentage)
	}

	// Test 7: Edge case - single element
	fmt.Println("\n=== Test 7: Single Element ===")
	rs7 := Constructor()
	fmt.Printf("insert(99): %v (expected: true)\n", rs7.Insert(99))
	fmt.Printf("getRandom(): %v (expected: 99)\n", rs7.GetRandom())
	fmt.Printf("getRandom(): %v (expected: 99)\n", rs7.GetRandom())
	fmt.Printf("remove(99): %v (expected: true)\n", rs7.Remove(99))

	// Test 8: Large operations
	fmt.Println("\n=== Test 8: Large Operations ===")
	rs8 := Constructor()

	// Insert 0-99
	for i := 0; i < 100; i++ {
		rs8.Insert(i)
	}
	fmt.Println("Inserted 100 elements (0-99)")

	// Remove even numbers
	for i := 0; i < 100; i += 2 {
		rs8.Remove(i)
	}
	fmt.Println("Removed 50 even numbers")

	// Check random values are only odd
	fmt.Print("10 random values (should all be odd): ")
	allOdd := true
	for i := 0; i < 10; i++ {
		val := rs8.GetRandom()
		fmt.Printf("%d ", val)
		if val%2 == 0 {
			allOdd = false
		}
	}
	fmt.Printf("\nAll odd? %v\n", allOdd)

	// Test 9: Alternating operations
	fmt.Println("\n=== Test 9: Alternating Operations ===")
	rs9 := Constructor()

	rs9.Insert(1)
	fmt.Println("insert(1): true")
	rs9.Insert(2)
	fmt.Println("insert(2): true")
	rs9.Remove(1)
	fmt.Println("remove(1): true")
	rs9.Insert(3)
	fmt.Println("insert(3): true")
	fmt.Printf("getRandom(): %v (expected: 2 or 3)\n", rs9.GetRandom())
	rs9.Remove(2)
	fmt.Println("remove(2): true")
	fmt.Printf("getRandom(): %v (expected: 3)\n", rs9.GetRandom())

	// Test 10: Negative numbers
	fmt.Println("\n=== Test 10: Negative Numbers ===")
	rs10 := Constructor()
	rs10.Insert(-5)
	rs10.Insert(-10)
	rs10.Insert(0)
	rs10.Insert(5)
	fmt.Println("Inserted: -5, -10, 0, 5")
	fmt.Printf("getRandom(): %v\n", rs10.GetRandom())
	fmt.Printf("remove(-10): %v (expected: true)\n", rs10.Remove(-10))
	fmt.Printf("remove(-10): %v (expected: false)\n", rs10.Remove(-10))
}

func main() {
	testRandomizedSet()
}
