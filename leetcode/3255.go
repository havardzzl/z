package leetcode

func resultsArray(nums []int, k int) []int {
	var vmax, curLen int
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] == vmax+1 {
			curLen++
		} else {
			curLen = 1
		}
		vmax = nums[i]
		if i < k-1 {
			continue
		}
		res := -1
		if curLen >= k {
			res = vmax
		}
		ans = append(ans, res)
	}
	return ans
}
