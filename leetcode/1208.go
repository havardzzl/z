package leetcode

import "math"

func equalSubstring(s string, t string, maxCost int) int {
	curCost := 0
	res := 0
	left, right := 0, 0
	n := len(s)
	getCost := func(i int) int {
		return int(math.Abs(float64(int(s[i]) - int(t[i]))))
	}
	for right < len(s) {
		cost := getCost(right)
		curCost += cost

		for curCost > maxCost && left <= right && left < n {
			curCost -= getCost(left)
			left++
		}
		if right-left+1 > res {
			res = right - left + 1
		}
		right++
	}
	return res
}
