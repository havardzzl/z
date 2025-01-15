package leetcode

import "math"

func maxPathSum(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	ans := math.MinInt
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		vLeft := dfs(node.Left)
		vRight := dfs(node.Right)
		ans = max(ans, node.Val+vLeft+vRight)
		return max(max(vLeft, vRight)+node.Val, 0)
	}
	dfs(root)
	return ans
}
