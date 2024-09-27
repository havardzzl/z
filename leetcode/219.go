package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	n := len(nums)
	for i := 0; i <= k && i < n; i++ {
		m[nums[i]]++
		if m[nums[i]] > 1 {
			return true
		}
	}
	for i := k + 1; i < n; i++ {
		m[nums[i-k-1]]--
		m[nums[i]]++
		if m[nums[i]] > 1 {
			return true
		}
	}
	return false
}
