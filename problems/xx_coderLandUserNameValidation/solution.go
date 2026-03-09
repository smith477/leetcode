package main

import "unicode"

func CodelandUsernameValidation(str string) bool {
	usernameLength := len(str)

	if usernameLength < 4 || usernameLength > 25 {
		return false
	}

	firstCharacter := rune(str[0])
	if !unicode.IsLetter(firstCharacter) {
		return false
	}

	lastCharacter := rune(str[usernameLength-1])
	if lastCharacter == '_' {
		return false
	}

	for _, character := range str {
		if !unicode.IsLetter(character) && !unicode.IsDigit(character) && character != '_' {
			return false
		}
	}

	return true
}
