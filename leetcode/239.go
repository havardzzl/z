package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	idxs := []int{}
	ans := []int{}
	for i := range nums {
		for len(idxs) > 0 && nums[idxs[len(idxs)-1]] <= nums[i] {
			idxs = idxs[:len(idxs)-1]
		}
		idxs = append(idxs, i)
		if i-idxs[0] >= k {
			idxs = idxs[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[idxs[0]])
		}
	}
	return ans
}
