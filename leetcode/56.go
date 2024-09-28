package leetcode

import (
	"sort"
)

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	ans := [][]int{}
	i := 0
	for i < len(intervals) {
		left := intervals[i][0]
		right := intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= right {
			right = max(right, intervals[j][1])
			j++
		}
		ans = append(ans, []int{left, right})
		i = j
	}
	return ans
}
