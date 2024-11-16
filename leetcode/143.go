package leetcode

func reorderList(head *ListNode) {
	node := head
	nodeStacks := []*ListNode{}
	for node != nil {
		nodeStacks = append(nodeStacks, node)
		node = node.Next
	}
	n := len(nodeStacks)
	left := head
	for i := n - 1; i > n-1-n/2; i-- {
		tmp := left.Next
		left.Next = nodeStacks[i]
		left.Next.Next = tmp
		left = tmp
	}
	left.Next = nil
}
