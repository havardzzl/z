package leetcode

import "sort"

func hIndex(citations []int) int {
	cs := sort.IntSlice(citations)
	n := len(cs)
	sort.Sort(cs)
	ans := 0
	for i, c := range cs {
		num := n - i
		if c >= num {
			ans = max(ans, num)
			break
		}
		ans = c
	}
	return ans
}
