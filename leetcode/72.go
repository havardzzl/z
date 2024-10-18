package leetcode

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = max(i, j)
			} else {
				v1 := dp[i-1][j-1]
				if word1[i-1] != word2[j-1] {
					v1++
				}
				v2 := min(dp[i-1][j], dp[i][j-1]) + 1
				dp[i][j] = min(v1, v2)
			}
		}
	}
	return dp[m][n]
}
