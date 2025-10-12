package leetcode

func maxArea(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	for l < r {
		ans = max(ans, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
