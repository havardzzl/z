package leetcode

import (
	"slices"
)

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	maxPint := 0
	var ans []int
	var dfs func(bobArrows []int, leftArrows int, index int, point int)
	dfs = func(bobArrows []int, leftArrows, index int, point int) {
		if index == 12 {
			if point > maxPint {
				ans = slices.Clone(bobArrows)
				maxPint = point
			}
			return
		}
		if index < 11 {
			dfs(bobArrows, leftArrows, index+1, point)
		}

		if index == 11 {
			bobArrows[index] = leftArrows
		} else if leftArrows > aliceArrows[index] {
			bobArrows[index] = aliceArrows[index] + 1
		}
		if bobArrows[index] > aliceArrows[index] {
			point += index
		}
		if index == 11 || bobArrows[index] > 0 {
			dfs(bobArrows, leftArrows-bobArrows[index], index+1, point)
			bobArrows[index] = 0
		}
	}
	dfs(make([]int, 12), numArrows, 0, 0)
	return ans
}
