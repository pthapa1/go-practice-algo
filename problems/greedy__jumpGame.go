package problems

import "math"

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

/*
function jump(nums: number[]): number {
  let jumpCount = 0;
  let currentEnd = 0;
  let farthest = 0;
  const lastIndex = nums.length - 1;
  for (let i = 0; i < lastIndex; i++) {
    farthest = Math.max(farthest, i + nums[i]);
    if (i === currentEnd) {
      jumpCount++;
      currentEnd = farthest;
      if (currentEnd >= lastIndex) break;
    }
  }
  return jumpCount;
}
*/
