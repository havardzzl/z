package leetcode

import (
	"math"
	"sort"
)

func nextPermutation(nums []int) {
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		v := math.MaxInt
		var pos int
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] && nums[j] < v {
				v = nums[j]
				pos = j
			}
		}
		if v < math.MaxInt {
			nums[i], nums[pos] = nums[pos], nums[i]
			sort.Ints(nums[i+1:])
			return
		}
	}
	sort.Ints(nums)
}
