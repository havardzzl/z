package leetcode

func integerReplacement(n int) int {
	steps := n
	var dfs func(num, step int)
	dfs = func(num, step int) {
		if num == 1 {
			steps = min(steps, step)
			return
		}
		if num%2 == 0 {
			dfs(num/2, step+1)
		} else {
			dfs(num-1, step+1)
			dfs(num+1, step+1)
		}
	}
	dfs(n, 0)
	return steps
}
