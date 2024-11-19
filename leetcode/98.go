package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	var isValid func(root *TreeNode, mi, ma int) bool
	isValid = func(root *TreeNode, mi, ma int) bool {
		if root == nil {
			return true
		}
		if root.Val <= mi || root.Val >= ma {
			return false
		}
		return isValid(root.Left, mi, root.Val) && isValid(root.Right, root.Val, ma)
	}
	return isValid(root, math.MinInt, math.MaxInt)
}
