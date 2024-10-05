package leetcode

func isIsomorphic(s string, t string) bool {
	m1 := map[byte]byte{}
	m2 := map[byte]struct{}{}
	for i := range s {
		m2[t[i]] = struct{}{}
		if m1[s[i]] == 0 {
			m1[s[i]] = t[i]
			continue
		}
		if m1[s[i]] != t[i] {
			return false
		}
	}
	return len(m1) == len(m2)
}
