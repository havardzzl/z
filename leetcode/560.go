package leetcode

func subarraySum(nums []int, k int) int {
	n := len(nums)
	sums := make([]int, n+1)
	for i := range nums {
		sums[i+1] = nums[i] + sums[i]
	}
	ans := 0
	cnt := map[int]int{0: 1}
	for i := 1; i < n+1; i++ {
		ans += cnt[sums[i]-k]
		cnt[sums[i]] += 1
	}
	return ans
}

/*
O(n**2)版本
func subarraySum(nums []int, k int) int {
	n := len(nums)
	// sums 保存以每个数字为起始位置的字数组和
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
*/
