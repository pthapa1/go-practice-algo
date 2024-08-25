package problems

// given a rotated array, find the minimum with 0(log n)
func RoatedArrayBinarySearch(arr []int) int {
	// it's a binary search so, it has pL, pR &&  calculate middleIndex -> middle number
	pL := 0
	pR := len(arr) - 1

	// we are trying to find out the pL, so for pL does not have to equal to < pR
	for pL < pR {
		middleIndex := (pL + pR) / 2
		if arr[pL] < arr[pR] {
			return arr[pL]
		}
		if arr[pR] < arr[middleIndex] {
			pL = middleIndex + 1
		} else {
			pR = middleIndex
		}
	}
	return arr[pL]
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func recurseRotatatedBs(arr []int, pL int, pR int) int {
	middleIndex := (pL + pR) / 2

	if pL == pR {
		return arr[pL]
	}
	// what if we can't decide.
	if arr[pL] == arr[pR] && arr[pR] == arr[middleIndex] {
		return mini(
			recurseRotatatedBs(arr, pL, middleIndex),
			recurseRotatatedBs(arr, middleIndex+1, pR),
		)
	}
	// non rotated array
	if arr[pL] < arr[pR] {
		return arr[pL]
	}
	// if the minimum is on the right side.
	if arr[pR] < arr[middleIndex] {
		return recurseRotatatedBs(arr, middleIndex+1, pR)
	}
	return recurseRotatatedBs(arr, pL, middleIndex)
}

func RotatedArrayBinarySearchWithRecursion(arr []int) int {
	return recurseRotatatedBs(arr, 0, len(arr)-1)
}
