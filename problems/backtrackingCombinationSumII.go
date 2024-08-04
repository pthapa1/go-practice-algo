/*
Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sum to target.
Each number in candidates may only be used once in the combination.
Note: The solution set must not contain duplicate combinations.

Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

Constraints:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/
package problems

import "slices"

func ComboSum2(candidates []int, target int) [][]int {
	var output [][]int
	var backtrack func(start int, target int, combination []int)
	backtrack = func(start int, target int, combination []int) {
		slices.Sort(candidates)
		if target < 0 {
			return
		}
		if target == 0 {
			aCopy := append([]int{}, combination...)
			output = append(output, aCopy)
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > target {
				break
			}
			combination = append(combination, candidates[i])
			backtrack(i+1, target-candidates[i], combination)
			combination = combination[:len(combination)-1]
		}
	}
	var combo []int
	backtrack(0, target, combo)
	if output == nil {
		return [][]int{}
	}
	return output
}
