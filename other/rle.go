package other

import (
	"strconv"
	"unicode"
)

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
	}
	return true
}

func decodeRle(input string) string {
	if len(input) == 0 {
		return input
	}

	if len(input) == 1 {
		return input
	}

	if !isASCII(input) {
		return input
	}

	output := ""
	multiplyCount := 1
	var prevChar rune

	for _, char := range input {
		if string(char) == "0" || isDigit(char) && isDigit(prevChar) {
			return input
		}
		if isDigit(char) {
			multiplyCount, _ = strconv.Atoi(string(char))
		} else {
			for i := 0; i < multiplyCount; i++ {
				output += string(char)
			}
		}
		prevChar = char
	}

	return output
}

func isValid(input string) bool {
	isValid := false
	for _, char := range input {
		isValid = (char >= 65 && char <= 90) || (char >= 97 && char <= 122)
		if !isValid {
			return false
		}
	}
	return isValid
}

func rle(input string) string {

	if len(input) == 0 {
		return ""
	}

	if !isValid(input) {
		return input
	}

	if len(input) == 1 {
		return input
	}

	const maxCharacterCount = 9
	output := ""
	prevChar := string(input[0])
	currentCount := 1

	for i := 1; i < len(input); i++ {

		currentChar := string(input[i])
		// for loop gives you byte. Convert to rune
		// range gives you rune.
		if currentChar == prevChar && currentCount < maxCharacterCount {
			currentCount++
		} else {
			output += strconv.Itoa(currentCount) + prevChar
			prevChar = currentChar
			currentCount = 1
		}
	}

	output += strconv.Itoa(currentCount) + prevChar

	return output
}
