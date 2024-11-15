package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur, next := head, head.Next
	cur.Next = nil
	for next != nil {
		cur, next, next.Next = next, next.Next, cur
	}
	return cur
}
