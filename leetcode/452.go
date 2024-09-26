package leetcode

import "sort"

type balloon struct {
	x     []int
	burst bool
}

type balloons []balloon

func (a balloons) Len() int           { return len(a) }
func (a balloons) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a balloons) Less(i, j int) bool { return a[i].x[1] < a[j].x[1] }

type indexStart struct {
	index int
	start int
}

type indexStarts []indexStart

func (a indexStarts) Len() int           { return len(a) }
func (a indexStarts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a indexStarts) Less(i, j int) bool { return a[i].start < a[j].start }

func FindMinArrowShots(points [][]int) int {
	bs := make(balloons, len(points))
	for i := range points {
		bs[i] = balloon{
			x: points[i],
		}
	}
	sort.Sort(bs)

	ans := 0
	starts := make(indexStarts, len(bs))
	for i, b := range bs {
		starts = append(starts, indexStart{
			index: i,
			start: b.x[0],
		})
	}
	sort.Sort(starts)
	for _, b := range bs {
		if b.burst {
			continue
		}
		b.burst = true
		idx, find := sort.Find(len(starts), func(i int) int {
			return b.x[1] - starts[i].start
		})
		for j := 0; j < idx; j++ {
			bs[starts[j].index].burst = true
		}
		if find {
			for idx < len(starts) && starts[idx].start == b.x[1] {
				bs[starts[idx].index].burst = true
				idx++
			}
		}
		ans++
	}
	return ans
}
