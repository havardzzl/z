package leetcode

import "strconv"

func letterCombinations(digits string) []string {
	m := map[byte][]byte{}
	for i := 2; i < 7; i++ {
		first := byte('a' + (i-2)*3)
		m[byte(strconv.Itoa(i)[0])] = []byte{first, first + 1, first + 2}
	}
	m['7'] = []byte{'p', 'q', 'r', 's'}
	m['8'] = []byte{'t', 'u', 'v'}
	m['9'] = []byte{'w', 'x', 'y', 'z'}
	n := len(digits)
	ans := []string{}
	if n == 0 {
		return ans
	}
	var dfs func(pos int, bs []byte)
	dfs = func(pos int, bs []byte) {
		if pos == n {
			ans = append(ans, string(bs))
			return
		}
		for _, c := range m[digits[pos]] {
			dfs(pos+1, append(bs, c))
		}
	}
	dfs(0, []byte{})
	return ans
}
