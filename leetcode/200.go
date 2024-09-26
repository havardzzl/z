package leetcode

// 会超时
func _overtime_numIslands(grid [][]byte) int {
	parent := map[int]int{}
	rank := map[int]int{}
	islandPos := []int{}
	for i, pots := range grid {
		for j, pot := range pots {
			pos := i*len(grid[0]) + j
			if pot == '1' {
				parent[pos] = pos
				rank[pos] = 0
				islandPos = append(islandPos, pos)
			}
		}
	}
	var find func(pos int) int
	find = func(pos int) int {
		if parent[pos] != pos {
			parent[pos] = find(parent[pos])
		}
		return parent[pos]
	}
	var union func(pos1, pos2 int) = func(pos1, pos2 int) {
		root1 := find(pos1)
		root2 := find(pos2)
		if root1 == root2 {
			return
		}
		if rank[root1] < rank[root2] {
			parent[root1] = root2
		} else if rank[root1] > rank[root2] {
			parent[root2] = root1
		} else {
			parent[root1] = root2
			rank[root2]++
		}
	}
	for i, pos1 := range islandPos {
		for _, pos2 := range islandPos[i+1:] {
			if (pos1/len(grid[0]) == pos2/len(grid[0]) && pos1+1 == pos2) || pos1+len(grid[0]) == pos2 {
				union(pos1, pos2)
			}
		}
	}
	for pos := range parent {
		find(pos)
	}
	rootIslands := map[int]bool{}
	for _, p := range parent {
		rootIslands[p] = true
	}
	return len(rootIslands)
}

func numIslands(grid [][]byte) int {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i >= len(grid) || i < 0 || j >= len(grid[0]) || j < 0 {
			return
		}
		if grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	ans := 0
	for i, pots := range grid {
		for j, pot := range pots {
			if pot == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}
