package leetcode

func getAverages(nums []int, k int) []int {
	res := make([]int, len(nums))
	s := 0
	n := len(nums)
	for i := 0; i < n && i < k; i++ {
		s += nums[i]
	}
	for i := 0; i < n; i++ {
		if i-k > 0 {
			s -= nums[i-k-1]
		}
		if i+k < n {
			s += nums[i+k]
		}
		if i-k < 0 || i+k >= n {
			res[i] = -1
		} else {
			res[i] = s / (k*2 + 1)
		}
	}
	return res
}
