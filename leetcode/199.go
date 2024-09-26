package leetcode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	nodes := []*TreeNode{root}
	goDown := func() {
		newNodes := []*TreeNode{}
		for i, node := range nodes {
			if i == 0 {
				ans = append(ans, node.Val)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
		}
		nodes = newNodes
	}
	for len(nodes) > 0 {
		goDown()
	}
	return ans
}
