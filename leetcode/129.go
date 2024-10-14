package leetcode

func sumNumbers(root *TreeNode) int {
	var total int
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		accu := sum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			total += accu
			return
		}
		dfs(node.Left, accu)
		dfs(node.Right, accu)
	}
	dfs(root, 0)
	return total
}
