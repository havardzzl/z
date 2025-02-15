package leetcode

import (
	"testing"
	"time"
)

func nt(v int) *TreeNode {
	return &TreeNode{Val: v}
}

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
	// t.Log("shopping price:", shoppingOffers([]int{6, 2, 8, 6, 10, 5}, [][]int{{5, 1, 6, 2, 0, 2, 19}, {3, 3, 5, 3, 5, 2, 2}, {6, 0, 4, 3, 2, 0, 14}, {5, 5, 4, 1, 6, 3, 23}, {3, 0, 5, 2, 1, 5, 35}, {1, 5, 4, 3, 1, 2, 36}, {5, 3, 5, 4, 3, 0, 1}, {6, 6, 4, 2, 4, 1, 5}, {3, 3, 2, 6, 1, 0, 33}, {2, 5, 1, 2, 4, 6, 23}, {3, 6, 2, 6, 2, 6, 14}, {6, 6, 0, 3, 3, 4, 17}, {0, 4, 5, 3, 5, 0, 15}, {6, 1, 0, 6, 4, 0, 14}, {6, 4, 4, 3, 3, 5, 8}, {4, 2, 4, 3, 6, 2, 30}, {3, 4, 0, 3, 1, 4, 3}, {4, 2, 6, 3, 3, 4, 12}, {6, 4, 2, 5, 1, 5, 16}, {3, 1, 0, 0, 3, 2, 3}}, []int{2, 4, 5, 3, 6, 3}))
	// t.Log("shop:", shoppingOffers([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}))
	// t.Log(countKConstraintSubstrings2("010101", 1, [][]int{{0, 5}, {1, 4}, {2, 3}}))
	root := nt(10)
	t21 := nt(5)
	t22 := nt(-3)
	t31 := nt(3)
	t32 := nt(2)
	t33 := nt(11)
	t41 := nt(3)
	t42 := nt(-2)
	t43 := nt(1)
	root.Left = t21
	root.Right = t22
	t21.Left = t31
	t21.Right = t32
	t22.Right = t33
	t31.Left = t41
	t31.Right = t42
	t32.Right = t43
	t.Log("countKConstrai2: ", pathSum(root, 8))
	time.Sleep(time.Second)
}
