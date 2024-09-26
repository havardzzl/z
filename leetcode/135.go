package leetcode

/*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。


*/

func candy(ratings []int) int {
	n := len(ratings)
	if n == 1 {
		return 1
	}
	lowPlace := []int{}
	for i := 0; i < n; {
		j := i
		for j < n-1 && ratings[j] == ratings[j+1] {
			j++
		}
		var isLow bool
		if i == 0 {
			isLow = j < n-1 && ratings[j] < ratings[j+1]
		} else if j == n-1 {
			isLow = i > 0 && ratings[i-1] > ratings[i]
		} else {
			isLow = ratings[i-1] > ratings[i] && ratings[j+1] > ratings[j]
		}
		if isLow {
			lowPlace = append(lowPlace, i)
		}
		i = j + 1
	}
	if len(lowPlace) == 0 {
		return len(ratings)
	}
	lowPlace = append([]int{0}, append(lowPlace, n-1)...)
	candys := make([]int, n)
	var calCandyLeft = func(target int) int {
		var compute int
		if ratings[target] == ratings[target+1] {
			compute = 1
		} else {
			compute = candys[target+1] + 1
		}
		if candys[target] != 0 {
			return max(compute, candys[target])
		}
		return compute
	}
	var calCandyRight = func(target int) int {
		if ratings[target] == ratings[target-1] {
			return 1
		}
		return candys[target-1] + 1
	}
	for i := 1; i < len(lowPlace)-1; i++ {
		p := lowPlace[i]
		candys[p] = 1
		for j := p - 1; j >= lowPlace[i-1]; j-- {
			if ratings[j] < ratings[j+1] {
				break
			}
			candys[j] = calCandyLeft(j)
		}
		for j := p + 1; j <= lowPlace[i+1]; j++ {
			if ratings[j] < ratings[j-1] {
				break
			}
			candys[j] = calCandyRight(j)
		}
	}
	allCandy := 0
	for i := 0; i < len(candys); i++ {
		allCandy += candys[i]
	}
	return allCandy
}
