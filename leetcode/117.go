package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	nodes := []*Node{root}
	for len(nodes) > 0 {
		newNodes := []*Node{}
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
