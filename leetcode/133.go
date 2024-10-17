package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	track := map[int]*Node{}
	record := map[int]*Node{}
	nodes := []*Node{node}
	for len(nodes) > 0 {
		nodes2 := []*Node{}
		for _, n := range nodes {
			track[n.Val] = n
			if record[n.Val] == nil {
				record[n.Val] = &Node{Val: n.Val}
			}
			for _, nb := range n.Neighbors {
				if record[nb.Val] == nil {
					nodes2 = append(nodes2, nb)
				}
			}
		}
		nodes = nodes2
	}
	for v := range record {
		for _, nb := range track[v].Neighbors {
			record[v].Neighbors = append(record[v].Neighbors, record[nb.Val])
		}
	}
	return record[node.Val]
}
