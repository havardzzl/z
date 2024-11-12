package leetcode

func countKConstraintSubstrings(s string, k int) int {
	var v0, v1 int
	n := len(s)
	ans := 0
	start := 0
	for i := 0; i < n && start < n; i++ {
		if s[i] == '0' {
			v0++
		} else {
			v1++
		}
		if v0 <= k || v1 <= k {
			continue
		}
		for v0 > k && v1 > k {
			ans += (n - i)
			if s[start] == '0' {
				v0--
			} else {
				v1--
			}
			start++
		}
	}
	return n*(n+1)/2 - ans
}
