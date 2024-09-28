package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	m := map[int]*ListNode{}
	node := head
	n := 0
	for node != nil {
		m[n] = node
		n++
		node = node.Next
	}
	k = k % n
	if k == 0 {
		return head
	}
	m[n-1].Next = head
	m[n-k-1].Next = nil
	return m[n-k]
}
