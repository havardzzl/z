package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]*bool, m)
	for i := range dp {
		dp[i] = make([]*bool, n)
	}
	var search func(i, j int) bool
	search = func(i, j int) bool {
		if dp[i][j] != nil {
			return *dp[i][j]
		}
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			return false
		}

		var searchDown, searchRight bool
		if i < m-1 {
			searchDown = search(i+1, j)
			dp[i+1][j] = &searchDown
		}
		if j < n-1 {
			searchRight = search(i, j+1)
			dp[i][j+1] = &searchRight
		}
		found := searchDown || searchRight
		dp[i][j] = &found
		return found
	}
	return search(0, 0)
}
