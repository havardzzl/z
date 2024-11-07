package leetcode

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var v int
	for i := 0; i < n; i++ {
		v += grid[0][i]
		dp[0][i] = v
	}
	v = 0
	for i := 0; i < m; i++ {
		v += grid[i][0]
		dp[i][0] = v
	}
	k := max(m, n)
	for i := 1; i < k; i++ {
		if i < m {
			for j := 1; j < n; j++ {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
		if i < n {
			for j := 1; j < m; j++ {
				dp[j][i] = min(dp[j-1][i], dp[j][i-1]) + grid[j][i]
			}
		}
	}
	return dp[m-1][n-1]
}
