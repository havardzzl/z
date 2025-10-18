package leetcode

func lengthOfLongestSubstring2(s string) int {
	ans := 0
	l, r := 0, 0
	n := len(s)
	cnt := map[byte]bool{}
	for r < n {
		if cnt[s[r]] {
			ans = max(ans, r-l)
			for s[l] != s[r] {
				cnt[s[l]] = false
				l++
			}
			l++
		}
		cnt[s[r]] = true
		r++
	}
	return max(ans, r-l)
}
