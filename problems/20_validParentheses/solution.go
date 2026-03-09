package main

import "fmt"

//func isValid(s string) bool {
//	var stack []rune
//
//	pairs := map[rune]rune{
//		')': '(',
//		']': '[',
//		'}': '{',
//	}
//
//	for _, char := range s {
//		if open, isClosing := pairs[char]; isClosing {
//			if len(stack) == 0 || stack[len(stack)-1] != open {
//				return false
//			}
//
//			stack = stack[:len(stack)-1]
//		} else {
//			stack = append(stack, char)
//		}
//	}
//
//	return len(stack) == 0
//}

func isValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if len(stack) != 0 {
			char := stack[len(stack)-1]
			if isPair(char, s[i]) {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}

	return len(stack) == 0
}

func isPair(last, curr byte) bool {
	isBracket := last == '(' && curr == ')'
	isSqaredBracket := last == '[' && curr == ']'
	isCurrlyBracket := last == '{' && curr == '}'
	return isBracket || isSqaredBracket || isCurrlyBracket
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
	fmt.Println(isValid(""))       // true
	fmt.Println(isValid("("))      // false
}
