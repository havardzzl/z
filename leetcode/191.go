package leetcode

import "math/bits"

func hammingWeight(n int) int {
	return bits.OnesCount(uint(n))
}
