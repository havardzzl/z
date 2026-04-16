package leetcode

func maxSum(nums []int, m int, k int) int64 {
	var res, sum int64
	records := map[int]int{}
	for i := 0; i < len(nums); i++ {
		records[nums[i]]++
		sum += int64(nums[i])
		if i+1 >= k {
			if i-k >= 0 {
				records[nums[i-k]]--
				sum -= int64(nums[i-k])
				if records[nums[i-k]] == 0 {
					delete(records, nums[i-k])
				}
			}
			if len(records) >= m && sum > res {
				res = sum
			}
		}
	}
	return res
}
