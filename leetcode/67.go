package leetcode

import "strconv"

func addBinary(a string, b string) string {
	l1, l2 := len(a), len(b)
	lmax := max(l1, l2)
	ans := ""
	var carry int
	for i := 0; i < lmax; i++ {
		if i < l1 {
			carry += int(a[l1-1-i] - '0')
		}
		if i < l2 {
			carry += int(b[l2-1-i] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry == 1 {
		return "1" + ans
	}
	return ans
}
