package leetcode

func maxSubarrayLength(nums []int, k int) int {
	cnt := map[int]int{}
	left := 0
	ans := 0
	for right, num := range nums {
		cnt[num]++
		for cnt[num] > k {
			cnt[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
