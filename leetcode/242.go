package leetcode

func isAnagram(s string, t string) bool {
	m := map[byte]int{}
	for _, b := range s {
		m[byte(b)]++
	}
	for _, b := range t {
		if m[byte(b)] == 0 {
			return false
		}
		m[byte(b)]--
		if m[byte(b)] == 0 {
			delete(m, byte(b))
		}
	}
	return len(m) == 0
}
