package leetcode

func losingPlayer(x int, y int) string {
	t := min(x, y/4)
	if t%2 == 0 {
		return "Bob"
	}
	return "Alice"
}
