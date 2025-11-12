package main

import "fmt"

// func getNextValidIndex(str string, idx int) int {
// 	skip := 0
// 	for 0 <= idx {
// 		if str[idx] == '#' {
// 			skip++
// 			idx--
// 		} else if skip > 0 {
// 			skip--
// 			idx--
// 		} else {
// 			break
// 		}
// 	}

// 	return idx
// }

// func backspaceCompare(s string, t string) bool {

// 	if s == t {
// 		return true
// 	}

// 	i := len(s) - 1
// 	j := len(t) - 1

// 	for i >= 0 || j >= 0 {
// 		i = getNextValidIndex(s, i)
// 		j = getNextValidIndex(t, j)

// 		if i < 0 && j < 0 {
// 			return true
// 		}

// 		if i < 0 || j < 0 {
// 			return false
// 		}

// 		if s[i] != t[j] {
// 			return false
// 		}

// 		i--
// 		j--
// 	}

// 	return true
// }

func backspaceCompare(s string, t string) bool {
	return backspace(s) == backspace(t)
}

func backspace(s string) string {
	st := []rune{}
	for _, r := range s {
		if r == '#' {
			if len(st) == 0 {
				continue
			} else {
				st = st[:len(st)-1]
			}
		} else {
			st = append(st, r)
		}
	}
	return string(st)
}

func main() {
	tests := []struct {
		s, t   string
		output bool
	}{
		{"ab#c", "ad#f#g#g#c", true},   // "ac" vs "ac"
		{"a##c", "#a#c", true},         // "c" vs "c"
		{"bxj##tw", "bxo#j##tw", true}, // "tw" vs "tw"
		{"nzp#o#g", "b#nzp#o#g", true}, // "nog" vs "nog"
		{"a###b", "b", true},           // "b" vs "b"
		{"y#fo##f", "y#f#o##f", true},  // "f" vs "f"
	}

	for _, test := range tests {
		result := backspaceCompare(test.s, test.t)
		fmt.Printf("s: %s, t: %s -> %v (expected %v)\n", test.s, test.t, result, test.output)
	}
}
