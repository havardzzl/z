package leetcode

func getCeiling(c int) int {
	s := 1
	for s <= c {
		s *= 2
		c /= 2
	}
	return s
}

func judgeSquareSum(c int) bool {
	if c <= 2 {
		return true
	}
	i, j := 0, getCeiling(c)
	for i <= j {
		s := i*i + j*j
		switch {
		case s == c:
			return true
		case s < c:
			i++
		default:
			j--
		}
	}
	return false
}
