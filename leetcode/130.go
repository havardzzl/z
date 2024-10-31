package leetcode

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	newVisit := func() [][]bool {
		vi := make([][]bool, m)
		for i := 0; i < m; i++ {
			vi[i] = make([]bool, n)
		}
		return vi
	}
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(i, j int, fn func(int, int), visit [][]bool)
	dfs = func(i, j int, fn func(int, int), visit [][]bool) {
		fn(i, j)
		visit[i][j] = true
		for _, dir := range dirs {
			r, c := i+dir[0], j+dir[1]
			if 0 <= r && r < m && 0 <= c && c < n && board[r][c] == 'O' && !visit[r][c] {
				dfs(r, c, fn, visit)
			}
		}
	}
	visit := newVisit()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visit[i][j] {
				var inborder bool
				dfs(i, j, func(r, c int) {
					if r == 0 || c == 0 || r == m-1 || c == n-1 {
						inborder = true
					}
				}, visit)
				if !inborder {
					vi := newVisit()
					dfs(i, j, func(r, c int) {
						board[r][c] = 'X'
					}, vi)
				}
			}
		}
	}
}
