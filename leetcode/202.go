package leetcode

import "fmt"

func isHappy(n int) bool {
	record := map[int]bool{n: true}
	for {
		parts := []int{}
		fmt.Println("begin gen parts, n: ", n)
		for n > 9 {
			parts = append(parts, n%10)
			n /= 10
		}
		parts = append(parts, n)
		fmt.Println("parts: ", parts)
		n = 0
		for _, part := range parts {
			n += part * part
		}

		fmt.Println("parts: ", parts, ", sum: ", n)
		if n == 1 {
			return true
		}
		if record[n] {
			return false
		}
		record[n] = true
	}
}
