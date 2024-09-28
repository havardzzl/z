package leetcode

func longestConsecutive(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num] = 1
	}
	ans := 0
	for num := range m {
		if m[num-1] > 0 {
			continue
		}
		length := 0
		for m[num] > 0 {
			length++
			num++
		}
		ans = max(ans, length)
	}
	return ans
}
