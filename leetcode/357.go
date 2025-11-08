package leetcode

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	avail := map[int]bool{}
	for i := range 10 {
		avail[i] = true
	}
	ans := 1
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans++
			return
		}
		for j := range 10 {
			// end at this pos
			if avail[j+1] {
				ans++
			}
			if i < n-1 && avail[j] {
				avail[j] = false
				dfs(i + 1)
				avail[j] = true
			}
		}
	}
	dfs(0)
	return ans
}
