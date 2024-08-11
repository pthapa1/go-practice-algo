package problems

/*
Given an integer array nums of unique elements, return all possible
subsets
 (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]
*/

func Subset(distinctNumbers []int) [][]int {
	emptyArr := []int{}
	output := [][]int{emptyArr}
	var backtrack func(start int, aSubset []int)
	backtrack = func(start int, aSubset []int) {
		// base case
		if start > len(distinctNumbers) {
			return
		}
		for i := start; i < len(distinctNumbers); i++ {
			candidate := distinctNumbers[i]
			aSubset = append(aSubset, candidate)
			// make a copy of aSubset and then push
			aCopyOfSubset := make([]int, len(aSubset))
			copy(aCopyOfSubset, aSubset)
			output = append(output, aCopyOfSubset)
			// then backtrack with a number infront of you
			backtrack(i+1, aSubset)
			// clean up
			aSubset = aSubset[:len(aSubset)-1]
		}
	}
	backtrack(0, []int{})
	return output
}
