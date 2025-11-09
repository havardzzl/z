package leetcode

import "slices"

func deleteAndEarn(nums []int) int {
	n := len(nums)
	slices.Sort(nums)
	getNextIdx := func(idx int) int {
		i := idx + 1
		for i < n && nums[i] == nums[idx] {
			i++
		}
		return i
	}
	f1, f2 := 0, 0
	f := 0
	lastV := 0
	i := 0
	for i < n {
		nextI := getNextIdx(i)
		earn := (nextI - i) * nums[i]
		if nums[i] > lastV+1 {
			f = f2 + earn
		} else {
			f = max(f2, f1+earn)
		}
		lastV = nums[i]
		f1, f2 = f2, f
		i = nextI
	}
	return f
}
