package leetcode

func gameOfLife(board [][]int) {
	// 1>0:3 0>1:2
	m, n := len(board), len(board[0])
	get := func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		return board[i][j]
	}
	dirs := [][2]int{}
	for i := -1; i < 2; i++ {
		if i == 0 {
			dirs = append(dirs, [2]int{i, -1}, [2]int{i, 1})
		} else {
			dirs = append(dirs, [2]int{i, -1}, [2]int{i, 0}, [2]int{i, 1})
		}
	}
	getSum := func(i, j int) int {
		ans := 0
		for _, dir := range dirs {
			v := get(i+dir[0], j+dir[1])
			ans += v
			if v > 1 {
				ans -= 2
			}
		}
		return ans
	}
	for i := range board {
		for j := range board[i] {
			cur := get(i, j)
			if cur > 1 {
				cur -= 2
			}
			sum := getSum(i, j)
			if cur == 1 && (sum < 2 || sum > 3) {
				cur = 3
			}
			if cur == 0 && sum == 3 {
				cur = 2
			}
			board[i][j] = cur
		}
	}
	for i := range board {
		for j := range board[i] {
			if v := get(i, j); v > 1 {
				board[i][j] = (v - 1) % 2
			}
		}
	}
}
