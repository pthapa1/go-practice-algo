package problems

import (
	"strings"
)

// given a string, reverse it's character.
func ReverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	array := strings.Split(s, "")
	lastCharacter := array[len(array)-1]
	reverseRest := ReverseString(s[:len(s)-1])
	return lastCharacter + reverseRest
}

// given a string, find if there are same two characters in a row
func FindConcurrentChar(s string) []byte {
	charCount := make(map[byte]int)
	var repeatingChars []byte

	for i := 0; i < len(s); i++ {
		char := s[i]
		charCount[char]++
	}

	for char, count := range charCount {
		if count > 1 {
			repeatingChars = append(repeatingChars, char)
		}
	}
	return repeatingChars
}
