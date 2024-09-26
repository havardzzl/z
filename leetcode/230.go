package leetcode

func kthSmallest(root *TreeNode, k int) int {
	var inOrder func(node *TreeNode)
	var ans int
	inOrder = func(node *TreeNode) {
		if node == nil || k == 0 {
			return
		}
		inOrder(node.Left)
		k--
		if k == 0 {
			ans = node.Val
		}
		inOrder(node.Right)
	}
	inOrder(root)
	return ans
}
