package problems

import (
	"strings"
)

// given a string, reverse it's character.
func ReverseString(s string) string {
	// hello should be olleh

	if len(s) <= 1 {
		return s
	}
	array := strings.Split(s, "")

	lastIndex := len(array) - 1
	lastCharacter := array[lastIndex]

	restReversed := ReverseString(s[:len(s)-1])

	return lastCharacter + restReversed
}
