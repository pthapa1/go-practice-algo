// remember, when you have sorted array, you have new powers you can utilize to solve your problems.
package algo

import (
	"fmt"
	"math"
)

func binarySearchList(haystack []int, needle int) bool {
	startingIndex := 0
	endingIndex := len(haystack) - 1

	for startingIndex <= endingIndex {
		var middleIndexFloat float64 = float64(startingIndex + ((endingIndex - startingIndex) / 2))
		middleIndex := int(math.Floor(middleIndexFloat))

		if needle < haystack[middleIndex] {
			endingIndex = middleIndex - 1
		}

		if needle > haystack[middleIndex] {
			startingIndex = startingIndex + 1
		}
		if needle == haystack[middleIndex] {
			return true
		}
	}
	return false
}

func ExecuteBinarySearchList() {
	// [1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420]
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	result := binarySearchList(foo, 69)
	fmt.Println(result)
}
