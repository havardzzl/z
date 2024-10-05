package leetcode

func majorityElement(nums []int) int {
	m := map[int]int{}
	var maxCnt, ans int
	for _, num := range nums {
		m[num]++
		if m[num] > maxCnt {
			maxCnt = m[num]
			ans = num
		}
	}
	return ans
}
