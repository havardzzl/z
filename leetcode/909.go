package leetcode

func getRowCol(idx, n int) (int, int) {
	b := idx - 1
	row := b / n
	if row%2 == 0 {
		return n - 1 - row, b % n
	}
	return n - 1 - row, n - 1 - b%n
}

func forward(queue *[]int, board [][]int, visit *[]int) {
	newQ := []int{}
	n := len(board)
	for _, start := range *queue {
		if start+6 >= n*n {
			(*visit)[n*n] = 1
			*queue = []int{}
			return
		}
		for step := 1; step <= 6; step++ {
			end := start + step
			row, col := getRowCol(end, n)
			if board[row][col] != -1 {
				end = board[row][col]
			}
			if (*visit)[end] == 1 {
				continue
			}
			(*visit)[end] = 1
			newQ = append(newQ, end)
		}
	}
	*queue = newQ
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	queue := []int{1}
	visit := make([]int, n*n+1)
	steps := 0
	for len(queue) > 0 && visit[n*n] != 1 {
		forward(&queue, board, &visit)
		steps++
	}
	if visit[n*n] == 1 {
		return steps
	}
	return -1
}
