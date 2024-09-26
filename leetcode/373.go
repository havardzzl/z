package leetcode

import (
	"container/heap"
)

type ints struct {
	idxs         [][]int
	nums1, nums2 []int
}

func (s ints) Len() int {
	return len(s.idxs)
}

func (s ints) Less(i, j int) bool {
	idxi := s.idxs[i]
	idxj := s.idxs[j]
	return s.nums1[idxi[0]]+s.nums2[idxi[1]] < s.nums1[idxj[0]]+s.nums2[idxj[1]]
}

func (s ints) Swap(i, j int) {
	s.idxs[i], s.idxs[j] = s.idxs[j], s.idxs[i]
}

func (s *ints) Push(x any) {
	s.idxs = append(s.idxs, x.([]int))
}

func (s *ints) Pop() any {
	obj := s.idxs[len(s.idxs)-1]
	s.idxs = s.idxs[:len(s.idxs)-1]
	return obj
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	hp := &ints{
		idxs:  [][]int{},
		nums1: nums1,
		nums2: nums2,
	}
	for i := 0; i < len(nums1); i++ {
		hp.idxs = append(hp.idxs, []int{i, 0})
	}
	ans := [][]int{}
	for hp.Len() > 0 && len(ans) < k {
		idx := heap.Pop(hp).([]int)
		ans = append(ans, []int{nums1[idx[0]], nums2[idx[1]]})
		if idx[1] < len(nums2)-1 {
			heap.Push(hp, []int{idx[0], idx[1] + 1})
		}
	}
	return ans
}
