package leetcode

import (
	"container/heap"
)

type ins []int

func (iv ins) Len() int {
	return len(iv)
}

func (iv ins) Less(i, j int) bool {
	return iv[i] > iv[j]
}

func (iv ins) Swap(i, j int) {
	iv[i], iv[j] = iv[j], iv[i]
}

func (iv *ins) Push(p any) {
	*iv = append(*iv, p.(int))
}

func (iv *ins) Pop() any {
	obj := (*iv)[len(*iv)-1]
	*iv = (*iv)[:len(*iv)-1]
	return obj
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	hp := &ins{}

	for i := 0; i < n-k+1; i++ {
		heap.Push(hp, nums[i])
	}
	tmpLarge := heap.Pop(hp).(int)
	for i := n - k + 1; i < n; i++ {
		if nums[i] >= tmpLarge {
			continue
		}
		heap.Push(hp, nums[i])
		tmpLarge = heap.Pop(hp).(int)
	}
	return tmpLarge
}
