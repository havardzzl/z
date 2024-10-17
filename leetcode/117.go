package leetcode

type Node11 struct {
	Val   int
	Left  *Node11
	Right *Node11
	Next  *Node11
}

func connect(root *Node11) *Node11 {
	if root == nil {
		return root
	}
	nodes := []*Node11{root}
	for len(nodes) > 0 {
		newNodes := []*Node11{}
		for i := range nodes {
			if i < len(nodes)-1 {
				nodes[i].Next = nodes[i+1]
			}
			if nodes[i].Left != nil {
				newNodes = append(newNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				newNodes = append(newNodes, nodes[i].Right)
			}
		}
		nodes = newNodes
	}
	return root
}
