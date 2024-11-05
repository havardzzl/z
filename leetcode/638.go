package leetcode

import (
	"fmt"
)

// 反面例子，会超时，没有记住中间状态，而且，special长度最多有50，遍历这种情况最多有50!，天文数据
// 但是可以看到 needs 长度只有6，大小只有10，遍历所有needs情况只有 (10+1)^6 种情况，小很多
func shoppingOffers__overtime__(price []int, special [][]int, needs []int) int {
	n := len(price)
	m := len(special)
	visit := make([]bool, m)

	getSpecialCnt := func(idx int) int {
		ans := 10
		for i := 0; i < n; i++ {
			if special[idx][i] != 0 {
				ans = min(ans, needs[i]/special[idx][i])
			}
		}
		return ans
	}
	subNeedsWithSpecial := func(idx, cnt int) {
		for i := 0; i < n; i++ {
			needs[i] -= special[idx][i] * cnt
		}
	}
	addNeedsWithSpecial := func(idx, cnt int) {
		for i := 0; i < n; i++ {
			needs[i] += special[idx][i] * cnt
		}
	}
	result := make([]int, m)
	var minPrice int
	for i := 0; i < n; i++ {
		minPrice += needs[i] * price[i]
	}
	var searchTimes int
	var search func(num int)
	search = func(num int) {
		if num == m {
			var curPrice int
			for i := 0; i < n; i++ {
				curPrice += needs[i] * price[i]
			}
			for i := 0; i < m; i++ {
				curPrice += result[i] * special[i][n]
			}
			minPrice = min(minPrice, curPrice)
			searchTimes++
			fmt.Println("cal price, result:", result, ",needs:", needs, "price:", curPrice, "search time:", searchTimes)
			return
		}
		for i := 0; i < m; i++ {
			if visit[i] {
				continue
			}
			cnt := getSpecialCnt(i)
			result[i] = cnt
			visit[i] = true
			subNeedsWithSpecial(i, cnt)
			search(num + 1)
			result[i] = 0
			visit[i] = false
			addNeedsWithSpecial(i, cnt)
		}
	}
	search(0)
	return minPrice
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	filterSpecial := [][]int{}
	for _, sp := range special {
		var normalPrice int
		for i := 0; i < n; i++ {
			normalPrice += sp[i] * price[i]
		}
		if sp[n] < normalPrice {
			filterSpecial = append(filterSpecial, sp)
		}
	}
	dp := map[string]int{}
	var dfs func(needBs []byte) int
	dfs = func(needBs []byte) int {
		if v, ok := dp[string(needBs)]; ok {
			return v
		}
		var minPrice int
		for i := 0; i < n; i++ {
			minPrice += int(needBs[i]) * price[i]
		}
	outer:
		for _, sp := range filterSpecial {
			leftNeeds := make([]byte, n)
			for i := 0; i < n; i++ {
				if needBs[i] < byte(sp[i]) {
					continue outer
				}
				leftNeeds[i] = needBs[i] - byte(sp[i])
			}

			dp[string(leftNeeds)] = dfs(leftNeeds)
			minPrice = min(minPrice, dp[string(leftNeeds)]+sp[n])
		}
		dp[string(needBs)] = minPrice
		return minPrice
	}
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = byte(needs[i])
	}
	return dfs(bs)
}
