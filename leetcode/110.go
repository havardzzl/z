package leetcode

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(getDepth(node.Left), getDepth(node.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ll := getDepth(root.Left)
	lr := getDepth(root.Right)
	if ll-lr > 1 || lr-ll > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
