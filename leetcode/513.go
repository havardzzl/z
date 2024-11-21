package leetcode

func findBottomLeftValue(root *TreeNode) int {
	var ans int
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		ans = nodes[0].Val
		nxtNodes := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				nxtNodes = append(nxtNodes, node.Left)
			}
			if node.Right != nil {
				nxtNodes = append(nxtNodes, node.Right)
			}
		}
		nodes = nxtNodes
	}
	return ans
}
