package leetcode

import (
	"math"
	"slices"
	"strconv"
)

func nextGreaterElement(n int) int {
	s := strconv.Itoa(n)
	l := len(s)
	chars := make([]byte, len(s))
	for i := range s {
		chars[i] = s[i]
	}
	maxStr := strconv.Itoa(math.MaxInt32)
	for i := l - 2; i >= 0; i-- {
		if chars[i] >= chars[i+1] {
			continue
		}
		for j := l - 1; j > i; j-- {
			if chars[j] > chars[i] {
				chars[i], chars[j] = chars[j], chars[i]
				slices.Reverse(chars[i+1:])
				resultS := string(chars)
				if len(resultS) > len(maxStr) || (len(resultS) == len(maxStr) && resultS > maxStr) {
					return -1
				}
				ans, _ := strconv.Atoi(resultS)
				return ans
			}
		}
	}
	return -1
}
