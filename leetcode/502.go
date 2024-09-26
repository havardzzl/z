package leetcode

import (
	"container/heap"
	"sort"
)

type invests []int

func (iv invests) Len() int {
	return len(iv)
}

func (iv invests) Less(i, j int) bool {
	return iv[i] > iv[j]
}

func (iv invests) Swap(i, j int) {
	iv[i], iv[j] = iv[j], iv[i]
}

func (iv *invests) Push(p any) {
	*iv = append(*iv, p.(int))
}

func (iv *invests) Pop() any {
	obj := (*iv)[len(*iv)-1]
	*iv = (*iv)[:len(*iv)-1]
	return obj
}

type pair struct {
	pro int
	cap int
}

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].cap < p[j].cap
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	ps := pairs{}
	for i := range profits {
		ps = append(ps, pair{
			pro: profits[i],
			cap: capital[i],
		})
	}
	sort.Sort(ps)
	iv := invests{}
	for k > 0 {
		k--
		idx := sort.Search(len(ps), func(i int) bool {
			return ps[i].cap > w
		})
		for _, p := range ps[:idx] {
			heap.Push(&iv, p.pro)
		}

		ps = ps[idx:]
		if iv.Len() == 0 {
			break
		}
		p := heap.Pop(&iv).(int)
		w += p
	}
	return w
}
