package leetcode

func findWinningPlayer(skills []int, k int) int {
	n := len(skills)
	var curWin, serieWins int
	for i := 1; i < n; i++ {
		if skills[curWin] > skills[i] {
			serieWins++
		} else {
			serieWins = 1
			curWin = i
		}
		if serieWins == k {
			return curWin
		}
	}
	return curWin
}
