package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	mp := map[byte]string{}
	ss := strings.Split(s, " ")
	if len(pattern) != len(ss) {
		return false
	}

	for i, sl := range ss {
		c := pattern[i] - 'a'
		if mp[c] != "" && mp[c] != sl {
			return false
		}
		mp[c] = sl
	}
	parts := map[string]struct{}{}
	for _, sl := range mp {
		parts[sl] = struct{}{}
	}
	return len(mp) == len(parts)
}
