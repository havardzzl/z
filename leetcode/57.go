package leetcode

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	start, end := newInterval[0], newInterval[1]
	n := len(intervals)
	i1 := sort.Search(n, func(i int) bool {
		return intervals[i][1] >= start
	})
	i2 := sort.Search(n, func(i int) bool {
		return intervals[i][1] >= end
	})
	if i1 == n {
		return append(intervals, newInterval)
	}
	newStart := min(start, intervals[i1][0])
	if i2 == n {
		return append(intervals[:i1], []int{newStart, end})
	}
	if intervals[i2][0] > end {
		return append(intervals[:i1], append([][]int{{newStart, end}}, intervals[i2:]...)...)
	}
	return append(intervals[:i1], append([][]int{{newStart, intervals[i2][1]}}, intervals[i2+1:]...)...)
}
