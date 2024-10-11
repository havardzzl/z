package leetcode

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	var up, left int
	down, right := m-1, n-1
	ans := make([]int, m*n)
	var step int
	for up <= down && left <= right {
		// up line
		for j := left; j <= right; j++ {
			ans[step+j-left] = matrix[up][j]
			fmt.Println("up: ", step+j-left, ans[step+j-left])
		}
		step += right - left + 1
		up++
		// right line
		for i := up; i < down; i++ {
			ans[step+i-up] = matrix[i][right]
			fmt.Println("right: ", step+i-up, ans[step+i-up])
		}
		if down > up {
			step += down - up
		}
		// bottom line
		if down >= up {
			for j := right; j >= left; j-- {
				ans[step+right-j] = matrix[down][j]
				fmt.Println("bot: ", step+right-j, ans[step+right-j])
			}
			step += right - left + 1
			right--
			down--
		}
		// left line
		if left < right {
			for i := down; i >= up; i-- {
				ans[step+down-i] = matrix[i][left]
				fmt.Println("left: ", step+down-i, ans[step+down-i])
			}
			if down >= up {
				step += down - up + 1
			}
			left++
		}
	}
	return ans
}
