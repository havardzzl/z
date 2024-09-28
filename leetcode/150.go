package leetcode

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	nums := []int{}
	cal := func(a, b int, s string) int {
		switch s {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		default:
			return a / b
		}
	}
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			n := len(nums)
			a, b := nums[n-2], nums[n-1]
			nums[n-2] = cal(a, b, token)
			nums = nums[:n-1]
		default:
			if num, err := strconv.Atoi(token); err == nil {
				nums = append(nums, num)
			}
		}
	}
	return nums[0]
}
