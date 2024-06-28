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

	function combinationSum(candidates: number[], target: number): number[][] {
	    const output: number[][] = []
		function backtrack(start, target, combination: number[] = []){
			// base case
			if(target <= 0){
				if(target === 0) output.push([...combination])
				return
			}
			for(let i=start;i<candidates.length;i++){
				const value = candidates[i]
				combination.push(value)
				backtrack(i, target - value, combination)
				combination.pop();
			}
		}
		backtrack(0, target,[])
		return output
	};

// Also, solve how you would do it if you can't repeat the same number twice
// this is how you pop the item from an array in go

	func RemoveIndex(s []int, index int) []int {
	    return append(s[:index], s[index+1:]...)
	}
*/
package problems
