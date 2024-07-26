/*
Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
Example 3:

Input: candidates = [2], target = 1
Output: []

Constraints:

1 <= candidates.length <= 30
2 <= candidates[i] <= 40
All elements of candidates are distinct.
1 <= target <= 40
*/
package problems

func ComboSum(choices []int, target int) [][]int {
	var output [][]int
	var backtrack func(start int, target int, aCombo []int)
	backtrack = func(start int, target int, aCombination []int) {
		if target <= 0 {
			if target == 0 {
				aCopy := append([]int{}, aCombination...)
				output = append(output, aCopy)
			}
			return
		}
		for i := start; i < len(choices); i++ {
			choice := choices[i]
			aCombination = append(aCombination, choice)
			backtrack(i, target-choice, aCombination)
			// pop the last item, weather its done or it's -ve
			// we need to pop the item of existing slice, (aCombination) without creating a new one
			aCombination = aCombination[:len(aCombination)-1]
		}
	}
	aCombination := []int{}
	backtrack(0, target, aCombination)
	if output == nil {
		return [][]int{}
	}
	return output
}
