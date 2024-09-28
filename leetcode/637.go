package leetcode

func averageOfLevels(root *TreeNode) []float64 {
	ans := []float64{}
	var cal func()
	nodes := []*TreeNode{root}
	cal = func() {
		if len(nodes) == 0 {
			return
		}
		var sum int
		var cnt int
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node != nil {
				sum += node.Val
				cnt++
				newNodes = append(newNodes, node.Left, node.Right)
			}
		}
		nodes = newNodes
		if cnt > 0 {
			ans = append(ans, float64(sum)/float64(cnt))
		}
		cal()
	}
	cal()
	return ans
}
