package leetcode

import (
	"testing"
)

func TestFn(t *testing.T) {
	// merge([]int{4, 0, 0, 0, 0, 0}, 1, []int{1, 2, 3, 5, 6}, 5)
	// t.Log("cnt: ", countByVal(45, 12, 312))
	// merge([]int{}, 0, []int{1}, 1)
	// t.Log("removeDuplicates: ", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	// t.Log("remove2Duplicates: ", remove2Duplicates([]int{1, 1, 1, 2, 2, 3}))
	// t.Log("candy: ", candy([]int{1, 2, 3, 5, 4, 3, 2, 1, 4, 3, 2, 1, 3, 2, 1, 1, 2, 3, 4}))
	// t.Log("isPalindrome: ", isPalindrome("A man, a plan, a canal: Panama"))
	// t.Log("isHappy: ", isHappy(7))
	// t.Log("numIslands: ", numIslands([][]byte{{'1', '1'}}))
	t.Log("isSymmetric:", isSymmetric(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}))
}
