package leetcode

func countCompleteDayPairs(hours []int) int {
	m := map[int]int{}
	for _, hour := range hours {
		m[hour%24]++
	}
	var ans int
	for h, num := range m {
		vh := (24 - h) % 24
		if m[h] <= 0 || m[vh] <= 0 {
			continue
		}
		cnt := num * m[vh]
		if h == 0 || h == 12 {
			cnt = num * (num - 1) / 2
		}
		ans += cnt
		m[h] -= num
		m[vh] -= num
	}
	return ans
}
