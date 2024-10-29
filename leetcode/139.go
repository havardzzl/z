package leetcode

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	wordM := map[string]bool{}
	for _, w := range wordDict {
		wordM[w] = true
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}
			if wordM[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
