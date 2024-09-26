package leetcode

import "sort"

func minimumTotal(triangle [][]int) int {
	records := []int{triangle[0][0]}
	for i := 1; i < len(triangle); i++ {
		nextLevelRecord := []int{
			records[0] + triangle[i][0],
		}
		for j := 1; j < i; j++ {
			nextLevelRecord = append(nextLevelRecord, min(records[j], records[j-1])+triangle[i][j])
		}
		nextLevelRecord = append(nextLevelRecord, records[len(records)-1]+triangle[i][i])
		records = nextLevelRecord
	}
	sted := sort.IntSlice(records)
	sort.Sort(sted)
	return sted[0]
}
