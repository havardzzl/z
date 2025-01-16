package leetcode

func moveZeroes(nums []int) {
	idxs := []int{}
	for i, num := range nums {
		if num == 0 {
			idxs = append(idxs, i)
		} else if len(idxs) > 0 {
			nums[idxs[0]] = num
			idxs = append(idxs[1:], i)
		}
	}
	n := len(nums)
	m := len(idxs)
	for i := 0; i < m; i++ {
		nums[n-1-i] = 0
	}
}
