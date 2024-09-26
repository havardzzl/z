// 模拟一个市场投资，每次投资有随机50%的投资者本金会翻倍，剩余50%的投资者本金变为原来1/3，输入本金和进行投资的人数，
// 进行模拟投资，投资回合数从m到n，每种回合数模拟k次，计算出该回合数下，投资结束时，所有人的资产平均值，中位数和90%分为数，
// 最后输出资产平均值最高的投资回合数，以及该回合数下，资产平均值，中位数和90%分为数

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	minRounds   = 5
	maxRounds   = 100
	simulations = 10
)

var (
	initialCapital float64 = 10
	numInvestors   int     = 1000
)

func simulateMarketInvestment(initialCapital float64, numInvestors int) {
	bestRounds := 0
	bestAvg, bestMedian, bestPercentile90 := 0.0, 0.0, 0.0

	for rounds := minRounds; rounds <= maxRounds; rounds++ {
		totalAvg := 0.0
		allResults := make([]float64, 0, numInvestors*simulations)

		for sim := 0; sim < simulations; sim++ {
			results := make([]float64, numInvestors)
			for i := range results {
				results[i] = initialCapital
			}

			for j := 0; j < rounds; j++ {
				indices := rand.Perm(numInvestors)
				for k, idx := range indices {
					if k < numInvestors/2 {
						results[idx] *= 1.5
					} else {
						results[idx] *= 0.6
					}
				}
			}

			for _, capital := range results {
				totalAvg += capital
				allResults = append(allResults, capital)
			}
		}

		avg := totalAvg / float64(numInvestors*simulations)
		sort.Float64s(allResults)
		median := allResults[len(allResults)/2]
		percentile90 := allResults[int(float64(len(allResults))*0.9)]

		fmt.Printf("%d次投资结束时，所有人的资产平均值,中位数,90%%分位数分别为：%.2f, %.2f, %.2f\n", rounds, avg, median, percentile90)

		if avg > bestAvg {
			bestRounds = rounds
			bestAvg = avg
			bestMedian = median
			bestPercentile90 = percentile90
		}
	}

	fmt.Printf("最佳投资回合数: %d\n", bestRounds)
	fmt.Printf("平均资产: %.2f\n", bestAvg)
	fmt.Printf("中位数资产: %.2f\n", bestMedian)
	fmt.Printf("90%%分位数资产: %.2f\n", bestPercentile90)
}

func main() {
	// fmt.Print("请输入初始本金: ")
	// fmt.Scan(&initialCapital)

	// fmt.Print("请输入投资人数: ")
	// fmt.Scan(&numInvestors)

	simulateMarketInvestment(initialCapital, numInvestors)
}
