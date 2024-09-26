package leetcode

import "fmt"

func remove2Duplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	notDupLen := 0
	i, j := 0, 1
	for j < len(nums)-1 {
		if nums[i] == nums[j] {
			if nums[j] != nums[j+1] {
				i, j = i+2, j+2
				notDupLen += 2
			} else {
				nums = append(nums[:j+1], nums[j+2:]...)
			}
		} else {
			i, j = i+1, j+1
			notDupLen += 1
		}
	}
	if j == len(nums)-1 {
		notDupLen += 2
	} else {
		notDupLen += 1
	}
	fmt.Println(nums[:notDupLen])
	return notDupLen
}
