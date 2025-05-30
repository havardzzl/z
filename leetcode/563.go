package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumL := dfs(node.Left)
		sumR := dfs(node.Right)
		ans += max(sumL, sumR) - min(sumL, sumR)
		return sumL + sumR + node.Val
	}
	dfs(root)
	return ans
}
