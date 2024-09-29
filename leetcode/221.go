package leetcode

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	ans := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = int(matrix[i][0] - '0')
		ans = max(ans, dp[i][0])
	}
	for j := 1; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
		ans = max(ans, dp[0][j])
	}
	leftUpOnes := func(row, col int) (int, int) {
		left, up := 0, 0
		for j := col; j >= 0; j-- {
			if matrix[row][j] == '0' {
				break
			}
			left++
		}
		for i := row; i >= 0; i-- {
			if matrix[i][col] == '0' {
				break
			}
			up++
		}
		return left, up
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			left, up := leftUpOnes(i, j)
			preValue := dp[i-1][j-1]
			curValue := 1
			switch {
			case left > preValue && up > preValue:
				curValue = preValue + 1
			default:
				curValue = min(left, up)
			}
			dp[i][j] = curValue
			ans = max(ans, dp[i][j])
		}
	}
	return ans * ans
}
