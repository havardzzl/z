package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var maxTime = 100

func printList(head *ListNode) {
	time := 0
	for head != nil {
		fmt.Printf("%d, ", head.Val)
		head = head.Next
		time++
		if time > maxTime {
			fmt.Print(".......")
			break
		}
	}
	fmt.Print("\n")
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	v := head.Val
	prev, cur := head, head.Next
	for cur != nil {
		if cur.Val == v {
			prev.Next = cur.Next
			cur = cur.Next
		} else {
			v = cur.Val
			prev, cur = cur, cur.Next
		}
	}
	return head
}
