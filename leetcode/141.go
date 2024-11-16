package leetcode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	a, b := head, head
	for b != nil && b.Next != nil {
		a, b = a.Next, b.Next.Next
		if a == b {
			return true
		}
	}
	return false
}
