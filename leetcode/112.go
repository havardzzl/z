package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	var found bool
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil || found {
			return
		}
		if node.Left == nil && node.Right == nil && sum+node.Val == targetSum {
			found = true
			return
		}
		dfs(node.Left, sum+node.Val)
		dfs(node.Right, sum+node.Val)

	}
	dfs(root, 0)
	return found
}
