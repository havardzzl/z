package leetcode

type NumArray struct {
	nums []int
	sums []int
}

func Constructor4(nums []int) NumArray {
	sums := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			sums[i] = nums[i]
		} else {
			sums[i] = nums[i] + sums[i-1]
		}
	}
	return NumArray{nums: nums, sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right] - this.sums[left] + this.nums[left]
}
