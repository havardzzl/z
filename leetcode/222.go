package leetcode

func countNodes(root *TreeNode) int {
	ans := 0
	nodes := []*TreeNode{root}
	base := 1
	for len(nodes) > 0 {
		if nodes[0] != nil {
			newNodes := []*TreeNode{}
			for _, node := range nodes {
				newNodes = append(newNodes, node.Right, node.Left)
			}
			nodes = newNodes
			ans += base

		} else {
			i := 1
			for i < len(nodes) {
				if nodes[i] != nil {
					break
				}
				i++
			}
			ans += base - i
			break
		}
		base *= 2
	}
	return ans
}
