package leetcode

func letterCasePermutation(s string) []string {
	switchCase := func(c rune) rune {
		if 'a' <= c && 'z' >= c {
			return c - 'a' + 'A'
		}
		return c - 'A' + 'a'
	}
	letters := []int{}
	ss := []rune{}
	for i, si := range s {
		ss = append(ss, si)
		if si >= '0' && si <= '9' {
			continue
		}
		letters = append(letters, i)
	}
	ans := []string{}
	var dfs func(rs []rune, pos int)
	dfs = func(rs []rune, pos int) {
		if pos == len(letters) {
			ans = append(ans, string(rs))
			return
		}
		dfs(rs, pos+1)
		rs[letters[pos]] = switchCase(rs[letters[pos]])
		dfs(rs, pos+1)
		rs[letters[pos]] = switchCase(rs[letters[pos]])
	}
	dfs(ss, 0)
	return ans
}
