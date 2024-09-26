package leetcode

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	steps := map[int]int{
		0: 1,
		1: 1,
	}
	i := 2
	for {
		steps[i] = steps[i-1] + steps[i-2]
		if i >= n {
			return steps[n]
		}
		i++
	}
}
