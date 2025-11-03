package leetcode

import (
	"slices"
)

func punished(n int) bool {
	nums := []int{}
	for k := n * n; k > 0; {
		nums = append(nums, k%10)
		k /= 10
	}
	slices.Reverse(nums)
	var flag bool
	var dfs func(i int, curSum, wholeSum int)
	dfs = func(i, curSum, wholeSum int) {
		if flag {
			return
		}
		if i == len(nums) {
			flag = curSum+wholeSum == n
			return
		}
		val := curSum*10 + nums[i]
		dfs(i+1, val, wholeSum)
		dfs(i+1, 0, wholeSum+val)
	}
	dfs(0, 0, 0)
	return flag
}

func punishmentNumber(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if punished(i) {
			sum += i * i
		}
	}
	return sum
}
