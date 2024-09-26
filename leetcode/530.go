package leetcode

import "math"

func getMinimumDifference(root *TreeNode) int {
	minDif := math.MaxInt
	var goThrough func(node *TreeNode) (mi, ma int)
	goThrough = func(node *TreeNode) (mi, ma int) {
		mi, ma = node.Val, node.Val
		if node.Left != nil {
			lmi, lma := goThrough(node.Left)
			minDif = min(minDif, node.Val-lma)
			mi = lmi
		}
		if node.Right != nil {
			rmi, rma := goThrough(node.Right)
			minDif = min(minDif, rmi-node.Val)
			ma = rma
		}
		return
	}
	goThrough(root)
	return minDif
}
