package leetcode

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (ans int) {
		defer func() {
			dp[i][j] = ans
		}()
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}
		if s[i] == s[j] {
			return 2 + dfs(i+1, j-1)
		}
		return max(dfs(i+1, j), dfs(i, j-1))
	}
	return dfs(0, n-1)
}
