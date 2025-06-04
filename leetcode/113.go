package leetcode

import "slices"

func pathSumV2(root *TreeNode, targetSum int) [][]int {
	paths := [][]int{}
	var dfs func(node *TreeNode, path []int, curSum int)
	dfs = func(node *TreeNode, path []int, curSum int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if node.Val+curSum == targetSum {
				paths = append(paths, slices.Clone(append(path, node.Val)))
			}
			return
		}
		dfs(node.Left, append(path, node.Val), curSum+node.Val)
		dfs(node.Right, append(path, node.Val), curSum+node.Val)
	}
	dfs(root, []int{}, 0)
	return paths
}
