package leetcode

// 定义线段树结构
type SegmentTree struct {
	tree []int
	size int
}

// 创建一个新的线段树
func NewSegmentTree(size int) *SegmentTree {
	return &SegmentTree{
		tree: make([]int, 4*size),
		size: size,
	}
}

// 更新线段树节点
func (st *SegmentTree) Update(index, value int) {
	st.update(1, 0, st.size-1, index, value)
}

func (st *SegmentTree) update(node, start, end, index, value int) {
	if start == end {
		// 叶节点，更新值
		st.tree[node] = max(st.tree[node], value)
	} else {
		mid := (start + end) / 2
		if index <= mid {
			st.update(2*node, start, mid, index, value)
		} else {
			st.update(2*node+1, mid+1, end, index, value)
		}
		// 更新当前节点的最大值
		st.tree[node] = max(st.tree[2*node], st.tree[2*node+1])
	}
}

// 查询线段树的区间最大值
func (st *SegmentTree) Query(left, right int) int {
	return st.query(1, 0, st.size-1, left, right)
}

func (st *SegmentTree) query(node, start, end, left, right int) int {
	if right < start || left > end {
		// 不相交的区间，返回负无穷
		return 0
	}
	if left <= start && end <= right {
		// 完全包含的区间
		return st.tree[node]
	}
	// 部分相交，递归查询
	mid := (start + end) / 2
	leftQuery := st.query(2*node, start, mid, left, right)
	rightQuery := st.query(2*node+1, mid+1, end, left, right)
	return max(leftQuery, rightQuery)
}

// 求解 LCIS 的函数
func LCIS_SegmentTree(arr1, arr2 []int) int {
	// 找出 arr2 的最大值，用于构建线段树
	maxVal := 0
	for _, v := range arr2 {
		if v > maxVal {
			maxVal = v
		}
	}

	// 构建线段树
	segTree := NewSegmentTree(maxVal + 1)

	// 遍历 arr1 和 arr2，利用动态规划和线段树更新 dp 值
	for _, num1 := range arr1 {
		// 每次遍历 num1 时，重置 maxDp
		maxDp := 0
		for _, num2 := range arr2 {
			if num1 == num2 {
				// 找到相等的元素，更新 dp
				newDp := maxDp + 1
				segTree.Update(num2, newDp)
			} else if num1 > num2 {
				// 如果 num1 > num2，更新 maxDp，查找以 arr2[j] 之前为结尾的最大 dp 值
				maxDp = max(maxDp, segTree.Query(0, num2))
			}
		}
	}

	// 最后，返回线段树中的最大值
	return segTree.Query(0, maxVal)
}

func LCIS(arr1, arr2 []int) int {
	n, m := len(arr1), len(arr2)
	if n == 0 || m == 0 {
		return 0
	}

	// dp array to store the length of LCIS for each position in arr2
	dp := make([]int, m)
	result := 0

	// Iterate through each element of arr1
	for i := 0; i < n; i++ {
		maxLength := 0 // Store the max LCIS length for arr1[i]

		// Iterate through each element of arr2
		for j := 0; j < m; j++ {
			// If arr1[i] > arr2[j], update maxLength (we can't extend LCIS at this point)
			if arr1[i] > arr2[j] {
				maxLength = max(maxLength, dp[j])
			}

			// If arr1[i] == arr2[j], we can extend the LCIS ending at arr2[j]
			if arr1[i] == arr2[j] {
				dp[j] = maxLength + 1
			}

			// Update the result with the longest LCIS found so far
			result = max(result, dp[j])
		}
	}

	return result
}
