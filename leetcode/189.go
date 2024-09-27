package leetcode

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	// i -> (i+k)%n

	step := 0
	for i := 0; i < n; i++ {
		var tmp int = nums[i]
		idx := i
		j := (idx + k) % n
		for j != idx {
			nums[j], tmp = tmp, nums[j]
			step++
			if step >= n {
				break
			}
			j = (j + k) % n
		}
		nums[idx] = tmp
		step++
		if step >= n {
			break
		}
	}
}
