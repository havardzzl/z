package leetcode

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode

	merge := func(m1, m2 map[int]bool) map[int]bool {
		if m1 == nil {
			if m2 == nil {
				return map[int]bool{}
			}
			return m2
		}
		for k := range m2 {
			m1[k] = true
		}
		return m1
	}
	var dfs func(node *TreeNode) map[int]bool
	dfs = func(node *TreeNode) map[int]bool {
		if node == nil || ans != nil {
			return nil
		}

		sawLeft := dfs(node.Left)
		sawRight := dfs(node.Right)
		saw := merge(sawLeft, sawRight)
		saw[node.Val] = true
		if saw[p.Val] && saw[q.Val] && ans == nil {
			ans = node
		}
		return saw
	}
	dfs(root)
	return ans
}
