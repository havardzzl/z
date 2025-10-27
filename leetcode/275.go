package leetcode

import "sort"

func hIndex2(citations []int) int {
	return sort.Search(min(len(citations), citations[len(citations)-1]), func(cnt int) bool {
		cnt++
		num := 0
		for _, ci := range citations {
			if ci >= cnt {
				num++
			}
		}
		return num < cnt
	})
}
