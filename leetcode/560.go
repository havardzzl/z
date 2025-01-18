package leetcode

func subarraySum(nums []int, k int) int {
	n := len(nums)
	sums := make([]int, n)
	ans := 0
	for i := range nums {
		for j := 0; j <= i; j++ {
			sums[j] += nums[i]
			if sums[j] == k {
				ans++
			}
		}
	}
	return ans
}
