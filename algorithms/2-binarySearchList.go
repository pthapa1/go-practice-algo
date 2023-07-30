// remember, when you have sorted array, you have new powers you can utilize to solve your problems.
package algo

func BinarySearch(haystack []int, needle int) bool {
	startingIndex := 0
	endingIndex := len(haystack) - 1

	for startingIndex <= endingIndex {
		var middleIndexFloat float64 = float64(startingIndex + ((endingIndex - startingIndex) / 2))
		middleIndex := int(middleIndexFloat)

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
