package leetcode

// optimize: 内层循环重复计算了左侧和右侧的最大值，用两个数组提前算好每个i左侧、右侧最大值
func trap(height []int) int {
	ans := 0
	n := len(height)
	for i := 1; i < n-1; i++ {
		lMax, rMax := height[i], height[i]
		for j := 0; j < i; j++ {
			lMax = max(lMax, height[j])
		}
		for j := i + 1; j < n; j++ {
			rMax = max(rMax, height[j])
		}
		if lMax > height[i] && rMax > height[i] {
			ans += min(lMax, rMax) - height[i]
		}
	}
	return ans
}
