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

func Permute(arr []int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}
	output := [][]int{}
	var backtrack func(remArr []int, aPermute []int)
	backtrack = func(remArr []int, aPermute []int) {
		// base case and when do we return
		if len(remArr) == 0 {
			aCopy := []int{}
			aCopy = append(aCopy, aPermute...)
			output = append(output, aCopy)
			return
		}
		for i := 0; i < len(remArr); i++ {
			candidate := remArr[i]
			aPermute = append(aPermute, candidate)
			newArr := append([]int{}, remArr[0:i]...)
			newArr = append(newArr, remArr[i+1:]...)
			backtrack(newArr, aPermute)
			aPermute = aPermute[:len(aPermute)-1]
		}
	}

	backtrack(arr, []int{})
	if output == nil {
		return [][]int{}
	}
	return output
}

// other solution from the internet

// array with arrays of int
//
// func Permute(arr []int) [][]int {
// 	if len(arr) == 0 {
// 		return [][]int{}
// 	}
//
// 	// define base case
// 	var output [][]int
// 	if len(arr) == 1 {
// 		return [][]int{arr}
// 	}
//
// 	for i := range arr {
// 		currElem := arr[i]
// 		remArr := append([]int{}, arr[:i]...)
// 		remArr = append(remArr, arr[i+1:]...)
// 		remPermutaiton := Permute(remArr)
// 		for _, v := range remPermutaiton {
// 			newPermutation := append([]int{currElem}, v...)
// 			output = append(output, newPermutation)
// 		}
// 	}
//
// 	return output
// }
