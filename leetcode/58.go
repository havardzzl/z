package leetcode

func lengthOfLastWord(s string) int {
	bs := []byte(s)
	n := len(bs)
	i := n - 1
	for ; i >= 0; i-- {
		if bs[i] != ' ' {
			break
		}
	}
	j := i
	for ; j >= 0; j-- {
		if bs[j] == ' ' {
			break
		}
	}
	return i - j
}
