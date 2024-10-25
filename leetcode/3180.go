package leetcode

import "slices"

func maxTotalReward(rewardValues []int) int {
	slices.Sort(rewardValues)
	m := rewardValues[len(rewardValues)-1]
	dp := make([]bool, m*2)
	dp[0] = true
	for _, rv := range rewardValues {
		for v := rv; v < 2*rv; v++ {
			dp[v] = dp[v] || dp[v-rv]
		}
	}
	for i := m*2 - 1; i >= 0; i-- {
		if dp[i] {
			return i
		}
	}
	return 0
}
