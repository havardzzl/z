package leetcode

func plusOne(digits []int) []int {
	add, left := 0, 0
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i] + add
		if i == len(digits)-1 {
			n++
		}
		add, left = n/10, n%10
		digits[i] = left
		if add == 0 {
			break
		}
	}
	if add > 0 {
		return append([]int{add}, digits...)
	}
	return digits
}
