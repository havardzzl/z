package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	record := map[*ListNode]bool{}
	node := headA
	for node != nil {
		record[node] = true
		node = node.Next
	}
	node = headB
	for node != nil {
		if record[node] {
			return node
		}
		node = node.Next
	}
	return nil
}
