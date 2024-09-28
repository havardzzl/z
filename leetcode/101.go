package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var symmetric func(lNodes, rNodes []*TreeNode) bool
	symmetric = func(lNodes, rNodes []*TreeNode) bool {
		m, n := len(lNodes), len(rNodes)
		if m != n {
			return false
		}
		if m == 0 {
			return true
		}
		lNodes2 := []*TreeNode{}
		rNodes2 := []*TreeNode{}
		for i := 0; i < m; i++ {
			if lNodes[i] == nil && rNodes[i] == nil {
				continue
			}
			if lNodes[i] == nil || rNodes[i] == nil {
				return false
			}
			if lNodes[i].Val != rNodes[i].Val {
				return false
			}
			lNodes2 = append(lNodes2, lNodes[i].Left, lNodes[i].Right)
			rNodes2 = append(rNodes2, rNodes[i].Right, rNodes[i].Left)
		}
		return symmetric(lNodes2, rNodes2)
	}
	return symmetric([]*TreeNode{root.Left}, []*TreeNode{root.Right})
}
