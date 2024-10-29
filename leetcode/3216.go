package leetcode

func getSmallestString(s string) string {
	n := len(s)
	for i := 0; i < n-1; i++ {
		v1 := s[i] - '0'
		v2 := s[i+1] - '0'
		if v1 <= v2 || v1%2 != v2%2 {
			continue
		}
		return string(append([]byte(s[:i]), append([]byte{v2 + '0', v1 + '0'}, []byte(s[i+2:])...)...))
	}
	return s
}
