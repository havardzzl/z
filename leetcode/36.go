package leetcode

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rowsNum := make(map[byte]bool)
		columnsNum := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if rowsNum[board[i][j]] {
					return false
				}
				rowsNum[board[i][j]] = true
			}
			if board[j][i] != '.' {
				if columnsNum[board[j][i]] {
					return false
				}
				columnsNum[board[j][i]] = true
			}
		}
	}
	squareNum := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		squareNum[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			row, col := i/3, j/3
			if squareNum[row*3+col][board[i][j]] {
				return false
			}
			squareNum[row*3+col][board[i][j]] = true
		}
	}
	return true
}
