package problems

import "math"

// https://leetcode.com/problems/jump-game-ii/description/
func JumpGameII(nums []int) int {
	totalJumps := 0
	destination := len(nums) - 1
	lastJumpIndex := 0 // how far can you jump from current position
	coverageWindow := 0
	for i := 0; i < len(nums); i++ {
		if len(nums) == 1 {
			return 0
		}
		coverageWindow = int(math.Max(float64(coverageWindow), float64(nums[i]+i)))
		// If you are at a point, that you have explored all the options, and your current index == lastOneYouCould jump to
		if i == lastJumpIndex {
			totalJumps++
			lastJumpIndex = coverageWindow
		}
		if lastJumpIndex >= destination {
			break
		}
	}

	return totalJumps
}

// https://leetcode.com/problems/jump-game/description/
func JumpGame(nums []int) bool {
	if len(nums) == 1 {
		return true // already at the last index
	}
	farthestPoint := 0
	coverageWindow := 0
	for i := 0; i < len(nums); i++ {
		// update the coverageWindow to whatever is the maximum, current range or the new one we are on
		coverageWindow = int(math.Max(float64(coverageWindow), float64(i+nums[i])))
		// if we are at the max point for the window above, explore more options
		if i == farthestPoint {
			farthestPoint = coverageWindow
			if coverageWindow >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
