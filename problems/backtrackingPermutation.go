package problems

/*
array: [1, 2, 3]
permuation_number = array.length = 3 => 3 * 2 * 1  = 6
permutations: [1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]

	             [1, 2, 3]  -> Full Array.
	            /    |    \ -> Remove current element and create remaining array
	      [2, 3]   [1, 3]   [1, 2] -> Recurse until arr.length === 1
	    /     /     \    \   \    \
	 [3]   [2]      [3] [1]   [2] [1] -> Then add the current what's left

	[1, 2, 3][1, 3, 2][2, 1, 3][2, 3, 1][3, 1, 2][3, 2, 1] -> push to output
*/

// array with arrays of int

func Permute(arr []int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}

	// define base case
	var output [][]int
	if len(arr) == 1 {
		return [][]int{arr}
	}

	for i := range arr {
		currElem := arr[i]
		remArr := append([]int{}, arr[:i]...)
		remArr = append(remArr, arr[i+1:]...)
		remPermutaiton := Permute(remArr)
		for _, v := range remPermutaiton {
			newPermutation := append([]int{currElem}, v...)
			output = append(output, newPermutation)
		}
	}

	return output
}
