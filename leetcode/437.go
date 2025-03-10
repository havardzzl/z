package leetcode

func pathSum(root *TreeNode, targetSum int) int {
	var ans int
	cnt := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, s int) {
		if node == nil {
			return
		}
		s += node.Val
		ans += cnt[s-targetSum]
		cnt[s]++
		dfs(node.Left, s)
		dfs(node.Right, s)
		cnt[s]--
	}
	dfs(root, 0)
	return ans
}
