package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
	var leftSum func(node *TreeNode, flag bool) int
	leftSum = func(node *TreeNode, flag bool) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			if flag {
				return node.Val
			}
			return 0
		}
		return leftSum(node.Left, true) + leftSum(node.Right, false)
	}
	return leftSum(root, false)
}
