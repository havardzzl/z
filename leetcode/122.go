package leetcode

func maxProfit(prices []int) int {
	n := len(prices)
	holds := make([]int, n)
	frees := make([]int, n)
	holds[0] = -prices[0]
	frees[0] = 0
	for i := 1; i < n; i++ {
		holds[i] = max(frees[i-1]-prices[i], holds[i-1])
		frees[i] = max(frees[i-1], holds[i-1]+prices[i])
	}
	return frees[n-1]
}
