package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		paths[0][i] = 1
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		paths[i][0] = 1
	}
	k := max(m, n)
	for i := 1; i < k; i++ {
		if i < m {
			for j := 1; j < n; j++ {
				if obstacleGrid[i][j] != 1 {
					paths[i][j] = paths[i-1][j] + paths[i][j-1]
				}
			}
		}
		if i < n {
			for j := 1; j < m; j++ {
				if obstacleGrid[j][i] != 1 {
					paths[j][i] = paths[j-1][i] + paths[j][i-1]
				}
			}
		}
	}
	return paths[m-1][n-1]
}
