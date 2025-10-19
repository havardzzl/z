package leetcode

func countSubarrays(nums []int, k int64) int64 {
	l := 0
	var ans int64 = 0
	var sum int64 = 0
	for r, val := range nums {
		num := int64(val)
		if num > k {
			l = r + 1
			sum = 0
			continue
		}
		sum += num
		for sum*int64(r-l+1) >= k {
			sum -= int64(nums[l])
			l++
		}
		ans += int64(r - l + 1)
	}
	return ans
}
