package leetcode

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	satisfied := 0
	n := len(customers)
	totalSuppress := make([]int, n+1)
	maxSuppress := 0
	for i := range customers {
		if grumpy[i] == 0 {
			satisfied += customers[i]
			totalSuppress[i+1] = totalSuppress[i]
		} else {
			totalSuppress[i+1] = totalSuppress[i] + customers[i]
		}
		if i+1 >= minutes {
			maxSuppress = max(maxSuppress, totalSuppress[i+1]-totalSuppress[i+1-minutes])
		}
	}
	if n <= minutes {
		return satisfied + totalSuppress[n]
	}
	return satisfied + maxSuppress
}
