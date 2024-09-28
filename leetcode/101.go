package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lNodes := []*TreeNode{root.Left}
	rNodes := []*TreeNode{root.Right}
	symmetric := func() bool {
		m, n := len(lNodes), len(rNodes)
		if m != n {
			return false
		}
		lNodes2 := []*TreeNode{}
		rNodes2 := []*TreeNode{}
		for i := 0; i < m; i++ {
			if lNodes[i] == nil && rNodes[i] == nil {
				continue
			}
			if (lNodes[i] == nil && rNodes[i] != nil) || (lNodes[i] != nil && rNodes[i] == nil) {
				return false
			}
			if lNodes[i].Val != rNodes[i].Val {
				return false
			}
			lNodes2 = append(lNodes2, lNodes[i].Left, lNodes[i].Right)
			rNodes2 = append(rNodes2, rNodes[i].Right, rNodes[i].Left)
		}
		lNodes = lNodes2
		rNodes = rNodes2
		return true
	}
	for len(lNodes) > 0 || len(rNodes) > 0 {
		if !symmetric() {
			return false
		}
	}
	return true
}
