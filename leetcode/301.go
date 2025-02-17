package leetcode

func removeInvalidParentheses(s string) []string {
	ans := []string{}
	bs := []byte{}
	n := len(s)
	maxValidLen := 0

	var dfs func(i int, open int)
	dfs = func(i int, open int) {
		if i == n {
			if open == 0 {
				ans = append(ans, string(bs))
				maxValidLen = max(maxValidLen, len(bs))
			}
			return
		}
		c := s[i]
		isL := c == '('
		isR := c == ')'
		dfs(i+1, open)
		if isR && open == 0 {
			return
		}
		bs = append(bs, c)
		if isL {
			dfs(i+1, open+1)
		} else if isR {
			dfs(i+1, open-1)
		} else {
			dfs(i+1, open)
		}
		bs = bs[:len(bs)-1]
	}
	dfs(0, 0)
	newAns := []string{}
	record := map[string]bool{}
	for _, v := range ans {
		if len(v) == maxValidLen && !record[v] {
			newAns = append(newAns, v)
			record[v] = true
		}
	}
	return newAns
}
