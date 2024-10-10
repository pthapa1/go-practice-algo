package problems

import "strings"

func greedy_removeKDigits(num string, k int) string {
	// loop over the num string.
	mySlice := []int{}
	completedSice := []string{}
	count := k
	for i := 0; i < len(num); i++ {
		currItem := num[i]
		for count > 0 && len(mySlice) > 0 && mySlice[len(mySlice)-1] > int(currItem) {
			count--
			mySlice = mySlice[:len(mySlice)-1]
		}
		mySlice = append(mySlice, int(currItem))
		completedSice = append(completedSice, string(currItem))
		for count > 0 {
			count--
			mySlice = mySlice[:len(mySlice)-1]
		}
	}
	return strings.Join(completedSice, "")
}
