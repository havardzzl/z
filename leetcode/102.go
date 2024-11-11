package leetcode

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	goThrough := func(nodes []*TreeNode) []*TreeNode {
		vs := []int{}
		nextLevelNodes := []*TreeNode{}
		for _, node := range nodes {
			vs = append(vs, node.Val)
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		ans = append(ans, vs)
		return nextLevelNodes
	}
	ns := []*TreeNode{root}
	for len(ns) > 0 {
		ns = goThrough(ns)
	}
	return ans
}
