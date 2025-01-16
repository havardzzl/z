package leetcode

func rob2(root *TreeNode) int {
	// return 0: rob 1: not rob
	var dfs func(node *TreeNode) [2]int
	dfs = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}
		left, right := dfs(node.Left), dfs(node.Right)
		doRob := node.Val + left[1] + right[1]
		notRob := max(left[0], left[1]) + max(right[0], right[1])
		return [2]int{doRob, notRob}
	}
	res := dfs(root)
	return max(res[0], res[1])
}

/*

超时的代码:
func rob(root *TreeNode) int {
    var dfs func(node *TreeNode, doRob bool) int
    dfs = func(node *TreeNode, doRob bool) int {
        if node == nil {
            return 0
        }
        if doRob {
            return node.Val + dfs(node.Left, false) + dfs(node.Right, false)
        }
        return max(dfs(node.Left, true), dfs(node.Left, false)) +
               max(dfs(node.Right, true), dfs(node.Right, false))
    }
    return max(dfs(root, false), dfs(root, true))
}
优化的代码中，状态转移方程没变，通过把多次遍历的结果通过多个返回值返回的方式降低了时间复杂度（ 由O(2**n) --> O(n) )
*/
