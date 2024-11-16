package leetcode

func detectCycle(head *ListNode) *ListNode {
	cnt := map[*ListNode]int{}
	for head != nil {
		cnt[head]++
		if cnt[head] > 1 {
			return head
		}
		head = head.Next
	}
	return nil
}
