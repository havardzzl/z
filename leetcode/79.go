package leetcode

func validRowCol(m, n, row, col int) bool {
	return row >= 0 && row < m && col >= 0 && col < n
}

func search(board [][]byte, visit [][]bool, row, col int, word string) bool {
	if len(word) == 0 {
		return true
	}
	m, n := len(board), len(board[0])
	steps := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < len(steps); i++ {
		newRow := row + steps[i][0]
		newCol := col + steps[i][1]
		if validRowCol(m, n, newRow, newCol) && !visit[newRow][newCol] && board[newRow][newCol] == word[0] {
			visit[newRow][newCol] = true
			if search(board, visit, newRow, newCol, word[1:]) {
				return true
			}
			visit[newRow][newCol] = false
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if word[0] != board[row][col] {
				continue
			}
			visit[row][col] = true
			if search(board, visit, row, col, word[1:]) {
				return true
			}
			visit[row][col] = false
		}
	}
	return false
}
