package leetcode

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ts := []int{}
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(ts) > 0 && temperatures[ts[len(ts)-1]] <= temperatures[i] {
			ts = ts[:len(ts)-1]
		}
		if len(ts) > 0 {
			ans[i] = ts[len(ts)-1] - i
		}
		ts = append(ts, i)
	}
	return ans
}

// 暴力解法会超时
/*
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	indexes := []int{0}
	for i := 1; i < n; i++ {
		t := temperatures[i]
		outs := []int{}
		newIndexes := []int{i}
		for _, index := range indexes {
			if t > temperatures[index] {
				outs = append(outs, index)
			} else {
				newIndexes = append(newIndexes, index)
			}
		}
		indexes = newIndexes
		for _, idx := range outs {
			ans[idx] = i - idx
		}
	}
	return ans
}
*/
