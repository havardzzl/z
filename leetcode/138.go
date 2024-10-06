package leetcode

type NNode struct {
	Val    int
	Next   *NNode
	Random *NNode
}

func copyRandomList(head *NNode) *NNode {
	type pair struct {
		val   int
		index *int
	}
	record := map[*NNode]int{}
	copyRecord := map[int]*NNode{}
	node := head
	pairs := []pair{}
	index := 0
	for node != nil {
		record[node] = index
		index++
		node = node.Next
	}
	node = head
	copyDummy := &NNode{}
	prev := copyDummy
	index = 0
	for node != nil {
		p := pair{val: node.Val}
		if node.Random != nil {
			idx := record[node.Random]
			p.index = &idx
		}
		pairs = append(pairs, p)
		cur := &NNode{
			Val: p.val,
		}
		prev.Next = cur
		copyRecord[index] = cur
		prev = cur
		node = node.Next
		index++
	}

	node = copyDummy.Next
	for _, p := range pairs {
		if p.index != nil {
			node.Random = copyRecord[*p.index]
		}
		node = node.Next
	}
	return copyDummy.Next
}
