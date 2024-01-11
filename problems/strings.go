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
