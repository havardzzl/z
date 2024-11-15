package leetcode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	cnt := 1
	var prev *ListNode
	cur := head
	for cnt < left {
		prev, cur = cur, cur.Next
		cnt++
	}
	h1, h2 := prev, cur

	next := cur.Next
	for cnt < right {
		cur, next, next.Next = next, next.Next, cur
		cnt++
	}
	h2.Next = next
	if left != 1 {
		h1.Next = cur
		return head
	}
	return cur
}
