package leetcode

/*

前提： 总gas量=总cost量 ==》所有路段的总富裕+总亏损=0
假设 找到的 [xi,...xj] (i<=j) 这个区间是富余汽油最多的区间，富余量为 k1
如果存在  n (n>j) 导致车开到n的时候没有油了 ==》[xj,...xn]这段路的总亏损 k2 > k1 ,[xi,...xn]总体是亏损的
又因为前提, 所以 [x0,...xi]+[xn,...end]这段路是富余的, 与假设矛盾


*/

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	dif := make([]int, n)
	for i := range cost {
		// dif is revert
		dif[i] = gas[n-1-i] - cost[n-1-i]
	}
	dp := make([]int, n)
	dp[0] = dif[0]
	minDp := make([]int, n)
	minDp[0] = dif[0]
	maxIdx := 0
	maxLeft := dp[0]
	minLeft := dp[0]
	sumDif := dif[0]
	for i := 1; i < n; i++ {
		sumDif += dif[i]
		dp[i] = max(dp[i-1]+dif[i], dif[i])
		minDp[i] = min(minDp[i-1]+dif[i], dif[i])
		if dp[i] > maxLeft {
			maxLeft = dp[i]
			maxIdx = i
		}
		if minDp[i] < minLeft {
			minLeft = minDp[i]
		}
	}
	if sumDif < 0 {
		return -1
	}
	maxLeft2 := sumDif - minLeft
	if maxLeft > maxLeft2 {
		return n - 1 - maxIdx
	}

	mv := dif[0]
	difSum := dif[0]
	mvIdx := 0
	for i := 1; i < n; i++ {
		difSum += dif[i]
		if mv < difSum {
			mvIdx = i
			mv = difSum
		}
	}
	return n - 1 - mvIdx
}
