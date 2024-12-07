package leetcode

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n)
	var s int
	for _, num := range nums {
		s += num
	}
	if target > 0 {
		s += target
	} else {
		s -= target
	}
	for i := 0; i < n; i++ {
		dp[i] = make([]int, s*2+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var dfs func(idx, tar int) int
	dfs = func(idx, tar int) int {
		if dp[idx][tar+s] >= 0 {
			return dp[idx][tar+s]
		}
		if idx == 0 {
			cnt := 0
			if nums[idx] == tar {
				cnt++
			}
			if nums[idx] == -tar {
				cnt++
			}
			dp[idx][tar+s] = cnt
			return cnt
		}
		res := dfs(idx-1, tar-nums[idx]) + dfs(idx-1, tar+nums[idx])
		dp[idx][tar+s] = res
		return res
	}
	return dfs(n-1, target)
}
