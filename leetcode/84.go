package leetcode

func largestRectangleArea(heights []int) int {
	n := len(heights)
	lefts := make([]int, n)
	rights := make([]int, n)
	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && heights[i] <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			lefts[i] = st[len(st)-1]
		} else {
			lefts[i] = -1
		}
		st = append(st, i)
	}
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && heights[i] <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			rights[i] = st[len(st)-1]
		} else {
			rights[i] = n
		}
		st = append(st, i)
	}
	ans := 0
	for i := range heights {
		ans = max(ans, heights[i]*(rights[i]-lefts[i]-1))
	}
	return ans
}
