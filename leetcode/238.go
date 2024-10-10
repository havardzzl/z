package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)
	fwdProduct := make([]int, n)
	bkdProduct := make([]int, n)
	ans := make([]int, n)
	p1 := 1
	p2 := 1
	for i := range nums {
		fwdProduct[i] = p1 * nums[i]
		p1 *= nums[i]
		bkdProduct[i] = p2 * nums[n-1-i]
		p2 *= nums[n-1-i]
	}
	ans[0] = bkdProduct[n-2]
	for i := 1; i < n-1; i++ {
		ans[i] = fwdProduct[i-1] * bkdProduct[n-2-i]
	}
	ans[n-1] = fwdProduct[n-2]
	return ans
}
