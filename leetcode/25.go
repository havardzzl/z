package leetcode

func reverseK(head *ListNode, k int) *ListNode {
	cur := head
	var prev *ListNode
	t := 0
	for t < k && cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
		t++
	}
	head.Next = cur
	return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	total := 0
	node := head
	for node != nil {
		total++
		node = node.Next
	}
	times := total / k
	node = head
	var groupHead, groupTail *ListNode
	var newHead *ListNode
	for t := 0; t < times; t++ {
		groupHead = reverseK(node, k)
		if groupTail != nil {
			groupTail.Next = groupHead
		}
		groupTail = node
		node = node.Next
		if t == 0 {
			newHead = groupHead
		}
	}
	return newHead
}
