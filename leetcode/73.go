package leetcode

func setZeroes(matrix [][]int) {
	rows := map[int]struct{}{}
	cols := map[int]struct{}{}
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				rows[i] = struct{}{}
				cols[j] = struct{}{}
			}
		}
	}
	for r := range rows {
		for j := range matrix[0] {
			matrix[r][j] = 0
		}
	}
	for c := range cols {
		for i := range matrix {
			matrix[i][c] = 0
		}
	}
}
