package leetcode

func maxVowels(s string, k int) int {
	m := map[int]bool{}
	for _, v := range []int{'a', 'e', 'i', 'o', 'u'} {
		m[v] = true
	}
	maxV := 0
	curV := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if m[int(s[i])] {
			curV++
		}
		if i >= k && m[int(s[i-k])] {
			curV--
		}
		maxV = max(maxV, curV)
	}
	return maxV
}
