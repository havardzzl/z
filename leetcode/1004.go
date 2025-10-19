package leetcode

func longestOnes(nums []int, k int) int {
	ans := 0
	l := 0
	zeros := []int{}
	for r, num := range nums {
		if num == 0 {
			if k > 0 {
				zeros = append(zeros, r)
				if len(zeros) > k {
					l = zeros[0] + 1
					zeros = zeros[1:]
				}
			} else {
				l = r + 1
			}
		}
		ans = max(ans, r-l+1)
	}
	return ans
}
