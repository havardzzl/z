package leetcode

import "math"

func minCameraCover(root *TreeNode) int {
	// 返回值：在node处安装/不按照摄像头时需要的最少摄像头数量，不按照摄像头时node能否被摄像头覆盖
	var dfs func(node *TreeNode) ([2]int, bool)
	dfs = func(node *TreeNode) ([2]int, bool) {
		if node == nil {
			return [2]int{math.MaxInt16, 0}, true
		}
		l, okL := dfs(node.Left)
		r, okR := dfs(node.Right)
		val := min(l[0], l[1]) + min(r[0], r[1])
		install := 1 + val
		// 左右节点在不安装摄像头时都能被覆盖到
		if okL && okR {
			// 左右子节点至少有一个安装摄像头时，node能通过子节点被覆盖
			cover := l[0] < l[1] || r[0] < r[1]
			return [2]int{install, val}, cover
		}
		// 左右节点至少有一个节点如果不安装摄像头就不能被覆盖
		return [2]int{install, math.MaxInt16}, true
	}
	vals, ok := dfs(root)
	if ok {
		return min(vals[0], vals[1])
	}
	return vals[0]
}
