package problems

import (
	"strings"
)

// Problem 1: given a string, reverse it's character using recursion
func ReverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	array := strings.Split(s, "")
	lastCharacter := array[len(array)-1]
	reverseRest := ReverseString(s[:len(s)-1])
	return lastCharacter + reverseRest
}

/*
Problem 2: given a string, find if there are same two characters in a row
Always: safety first return if the string is empty

// initial stuff
-  Convert String to Lowercase:
-  Check other stuff as necessary, like regex patterns for other things you want to avoid

// Initialize Variables:
positionCharCount to hold the final result.
tempPcc to temporarily hold data for comparison.
Two pointers, pointer1 and pointer2, are initialized.
pointer1 starts at the beginning of the string, and pointer2 starts at the second character.
Iterate Through the String while pointer 2 is less than the length of the string since pointer is index based

Compare Characters:

If the characters at pointer1 and pointer2 are the same, Update tempPcc with
the current character, its starting index, and the count of how many times it has been repeated consecutively so far.
then, increment pointer 2 to check the next character.

Update Final Result:
After finding a sequence of repeating characters,
Check if this sequence is longer than any previously stored sequence
If it is, positionCharCount is updated with the information from tempPcc.

Handle Non-Repeating Characters:
If the characters at pointer1 and pointer2 are different,
pointer1 is moved to pointer2's position, and pointer2 is incremented.

Return the struct:
Solution: 👇
*/
type CharProtocol struct {
	Character     string
	StartingIndex int
	Count         int
}

func FindConcurrentChar(s string) CharProtocol {
	var positionCharCount CharProtocol //  or ==> positionCharCount := CharProtocol{}
	if len(s) == 0 {
		return positionCharCount
	}
	s = strings.ToLower(s)   // ignore case sensitivity
	var tempPcc CharProtocol // holds the temp struct to compare with what exists
	pointer1 := 0
	pointer2 := 1

	for pointer2 < len(s) { // pointer1 & pointer2 are index based while lengh is not
		itemAtPointer1 := s[pointer1]
		itemAtPointer2 := s[pointer2]

		// if the item is the same
		if itemAtPointer1 == itemAtPointer2 {
			// create a temp object
			tempPcc.StartingIndex = pointer1
			tempPcc.Character = string(itemAtPointer1)
			tempPcc.Count = pointer2 - pointer1 + 1 // count starts with 2, distance plus 1
			pointer2++
			// only update the return map if we find a sequence larger than what exists
			if positionCharCount.Count < tempPcc.Count {
				positionCharCount.StartingIndex = tempPcc.StartingIndex
				positionCharCount.Character = tempPcc.Character
				positionCharCount.Count = tempPcc.Count
			}
		}

		// if the item is different
		if itemAtPointer1 != itemAtPointer2 {
			// update both pointers
			pointer1 = pointer2
			pointer2++
		}
	}
	return positionCharCount
}
