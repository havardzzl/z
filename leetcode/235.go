package leetcode

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	pv, qv := p.Val, q.Val
	for root != nil {
		if pv < root.Val && qv < root.Val {
			root = root.Left
			continue
		}
		if pv > root.Val && qv > root.Val {
			root = root.Right
			continue
		}
		return root
	}
	return nil
}
