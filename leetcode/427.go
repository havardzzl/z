package leetcode

type NNode0 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *NNode0
	TopRight    *NNode0
	BottomLeft  *NNode0
	BottomRight *NNode0
}

func theSame(grid [][]int) bool {
	base := grid[0][0]
	for _, col := range grid {
		for _, bit := range col {
			if bit != base {
				return false
			}
		}
	}
	return true
}

func construct(grid [][]int) *NNode0 {
	if theSame(grid) {
		return &NNode0{
			Val:    grid[0][0] == 1,
			IsLeaf: true,
		}
	}
	n := len(grid)
	topL := [][]int{}
	topR := [][]int{}
	botL := [][]int{}
	botR := [][]int{}
	for i := 0; i < n/2; i++ {
		topL = append(topL, grid[i][:n/2])
		topR = append(topR, grid[i][n/2:])
		botL = append(botL, grid[n/2+i][:n/2])
		botR = append(botR, grid[n/2+i][n/2:])
	}
	return &NNode0{
		IsLeaf:      false,
		TopLeft:     construct(topL),
		TopRight:    construct(topR),
		BottomLeft:  construct(botL),
		BottomRight: construct(botR),
	}
}
