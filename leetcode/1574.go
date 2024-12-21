package leetcode

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	left, right := 0, n-1
	for right-1 >= 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 {
		return 0
	}
	i := 0
	for i+1 < n && arr[i] <= arr[i+1] {
		i++
	}
	ans := min(right, n-1-i)
	leftBound := i
	for right < n {
		for left <= leftBound && arr[left] <= arr[right] {
			left++
		}
		ans = min(ans, right-left)
		right++
	}

	return ans
}
