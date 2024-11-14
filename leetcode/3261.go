package leetcode

import (
	"sort"
)

func countKConstraintSubstrings2(s string, k int, queries [][]int) []int64 {
	n := len(s)
	cnt := make([]int, n)
	sum := make([]int, n+1)
	vs := [2]int{}
	var start int
	for i := 0; i < n; i++ {
		vs[s[i]-'0']++
		for vs[0] > k && vs[1] > k {
			vs[s[start]-'0']--
			cnt[start] = i - start
			sum[start+1] = sum[start] + cnt[start]
			start++
		}
	}
	for start < n {
		cnt[start] = n - start
		sum[start+1] = sum[start] + cnt[start]
		start++
	}
	ans := make([]int64, len(queries))
	for i, q := range queries {
		idx := sort.Search(q[1]-q[0]+1, func(i int) bool {
			i += q[0]
			return cnt[i]+i-1 >= q[1]
		})
		ans[i] = int64((q[1] - q[0] - idx + 1) * (q[1] - q[0] - idx + 2) / 2)
		ans[i] += int64(sum[q[0]+idx] - sum[q[0]])
	}
	return ans
}
