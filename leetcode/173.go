package leetcode

type BSTIterator struct {
	data []int
}

func Constructor(root *TreeNode) BSTIterator {
	var inorder func(node *TreeNode)
	data := []int{}
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		data = append(data, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return BSTIterator{
		data: data,
	}
}

func (this *BSTIterator) Next() int {
	v := this.data[0]
	this.data = this.data[1:]
	return v
}

func (this *BSTIterator) HasNext() bool {
	return len(this.data) > 0
}
