package leetcode

func countArrangement(n int) int {
	avail := map[int][]int{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				avail[i] = append(avail[i], j)
			}
		}
	}
	ans := 0
	occupy := map[int]bool{}
	var dfs func(k int)
	dfs = func(k int) {
		if k == n+1 {
			ans++
		}
		for _, v := range avail[k] {
			if occupy[v] {
				continue
			}
			occupy[v] = true
			dfs(k + 1)
			occupy[v] = false
		}
	}
	dfs(1)
	return ans
}
