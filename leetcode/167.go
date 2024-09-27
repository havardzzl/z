package leetcode

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		s := numbers[i] + numbers[j]
		if s == target {
			break
		}
		if s < target {
			i++
			continue
		}
		j--
	}
	return []int{i + 1, j + 1}
}
