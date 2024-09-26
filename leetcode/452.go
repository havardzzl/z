package leetcode

import "sort"

func FindMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	ans := 0
	arrowPos := points[0][1]
	for _, p := range points {
		if p[0] > arrowPos {
			ans++
			arrowPos = p[1]
		}
	}
	return ans
}
