package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	m := map[byte]int{}
	for _, c := range magazine {
		m[byte(c-'a')]++
	}
	for _, c := range ransomNote {
		m[byte(c-'a')]--
		if m[byte(c-'a')] < 0 {
			return false
		}
	}
	return true
}
