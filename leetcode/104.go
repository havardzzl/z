package leetcode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var bfs func(nodes []*TreeNode) []*TreeNode
	ans := 0
	bfs = func(nodes []*TreeNode) []*TreeNode {
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
		return newNodes
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		nodes = bfs(nodes)
		ans++
	}
	return ans
}
