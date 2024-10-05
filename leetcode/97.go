package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}
	if l3 == 0 {
		return true
	}
	dp := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]bool, l2+1)
	}

	dp[0][0] = true
	if l1 > 0 {
		dp[1][0] = s1[0] == s3[0]
	}
	if l2 > 0 {
		dp[0][1] = s2[0] == s3[0]
	}

	for i3 := 2; i3 <= l3; i3++ {
		for i1 := 0; i1 <= i3 && i1 <= l1; i1++ {
			i2 := i3 - i1
			if i2 > l2 {
				continue
			}
			if i1 > 0 && s1[i1-1] == s3[i3-1] {
				dp[i1][i2] = dp[i1][i2] || dp[i1-1][i2]
			}
			if i2 > 0 && s2[i2-1] == s3[i3-1] {
				dp[i1][i2] = dp[i1][i2] || dp[i1][i2-1]
			}
		}
	}
	return dp[l1][l2]
}
