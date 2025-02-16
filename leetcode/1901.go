package leetcode

func indexOfMax(a []int) (idx int) {
	for i, x := range a {
		if x > a[idx] {
			idx = i
		}
	}
	return
}

func findPeakGrid(mat [][]int) []int {
	m := len(mat)
	if m == 1 {
		return []int{0, indexOfMax(mat[0])}
	}
	l, r := 0, m-2
	for l <= r {
		mid := (l + r) / 2
		pos := indexOfMax(mat[mid])
		if mid > 0 && mat[mid-1][pos] > mat[mid][pos] {
			r = mid - 1
		} else if mat[mid][pos] < mat[mid+1][pos] {
			l = mid + 1
		} else {
			l = mid
			break
		}
	}
	return []int{l, indexOfMax(mat[l])}
}
