package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	var getMaxDepth func(node *TreeNode) int
	ans := 0
	getMaxDepth = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		leftDep := getMaxDepth(node.Left)
		rightDep := getMaxDepth(node.Right)
		ans = max(ans, leftDep+rightDep+2)
		return max(leftDep, rightDep) + 1
	}
	getMaxDepth(root)
	return ans
}
