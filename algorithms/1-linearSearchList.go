package algo

import (
	"fmt"
)

func linearSearch(haystack []int, needle int) bool {

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

// var foo = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420} /* Slice of Int */
var foo = [11]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420} /* Array of Int. Arrays have fixed length */

func ExecuteLinearSearch() {
	boolValue := linearSearch(foo[:], 0)
	fmt.Println(boolValue)
}
