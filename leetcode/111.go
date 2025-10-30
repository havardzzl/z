package leetcode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		if root.Right == nil {
			return 1
		}
		return minDepth(root.Right) + 1
	} else {
		if root.Right != nil {
			return min(minDepth(root.Left), minDepth(root.Right)) + 1
		}
		return minDepth(root.Left) + 1
	}
}
