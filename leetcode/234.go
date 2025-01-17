package leetcode

func isPalindrome3(head *ListNode) bool {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	n := len(vals)
	for i := 0; i < n/2; i++ {
		if vals[i] != vals[n-1-i] {
			return false
		}
	}
	return true
}
