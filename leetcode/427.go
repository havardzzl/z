package leetcode

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
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

func construct(grid [][]int) *Node {
	if theSame(grid) {
		return &Node{
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
	return &Node{
		IsLeaf:      false,
		TopLeft:     construct(topL),
		TopRight:    construct(topR),
		BottomLeft:  construct(botL),
		BottomRight: construct(botR),
	}
}
