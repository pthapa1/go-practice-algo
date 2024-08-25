package problems

func recurse(arr []int, target int, pL int, pR int) bool {
	if len(arr) == 0 {
		return false
	}
	middleIndex := (pL + pR) / 2
	middleNumber := arr[middleIndex]
	if pL > pR {
		return false
	}
	if target == middleNumber {
		return true
	}
	if target < middleNumber {
		return recurse(arr, target, pL, middleIndex-1)
	} else {
		return recurse(arr, target, middleIndex+1, pR)
	}
}

// Given a sorted array, return true if target exists. Make sure it's the O(log n)
func RecursiveBinarySearch(arr []int, target int) bool {
	return recurse(arr, target, 0, len(arr)-1)
}

func LoopBinarySearch(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}
	pL := 0
	pR := len(arr) - 1

	for pL <= pR {
		middleIndex := (pL + pR) / 2
		if target == arr[middleIndex] {
			return true
		}
		if target < arr[middleIndex] {
			pR = middleIndex - 1
		} else {
			pL = middleIndex + 1
		}
	}
	return false
}
