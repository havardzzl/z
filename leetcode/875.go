package leetcode

import "slices"

func upperDiv(a, b int) int {
	s := a / b
	if a%b == 0 {
		return s
	}
	return s + 1
}

func minEatingSpeed(piles []int, h int) int {
	slices.Sort(piles)
	l, r := upperDiv(piles[0], h), piles[len(piles)-1]
	for l <= r {
		mid := l + (r-l)/2
		total := 0
		for _, pi := range piles {
			total += pi / mid
			if pi%mid != 0 {
				total++
			}
		}
		if total > h {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
