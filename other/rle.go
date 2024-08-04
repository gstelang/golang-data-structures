package other

import (
	"strconv"
)

func rle(input string) string {

	if len(input) == 0 {
		return ""
	}

	if len(input) == 1 {
		return "1" + string(input[0])
	}

	const maxChars = 9
	output := ""
	prevChar := string(input[0])
	currentCount := 1

	for i := 1; i < len(input); i++ {
		currentChar := string(input[i])
		if currentChar == prevChar {
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
