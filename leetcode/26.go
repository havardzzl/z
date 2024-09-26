package leetcode

import "fmt"

func removeDuplicates(nums []int) int {
	notDup := 1
	for i := 0; i < len(nums)-1; {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i+1], nums[i+2:]...)
		} else {
			notDup++
			i++
		}
	}
	fmt.Println(nums[:notDup])
	return notDup
}
