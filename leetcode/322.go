package leetcode

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, amount+1)
	}
	var dfs func(idx, amt int) int
	dfs = func(idx, amt int) (res int) {
		if amt == 0 {
			return 0
		}
		if amt < 0 || idx < 0 {
			return amount + 1
		}
		if dp[idx][amt] > 0 {
			return dp[idx][amt]
		}
		defer func() { dp[idx][amt] = res }()
		return min(dfs(idx-1, amt), dfs(idx, amt-coins[idx])+1)
	}
	ans := dfs(n-1, amount)
	if ans > amount {
		return -1
	}
	return ans
}
