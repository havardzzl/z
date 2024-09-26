package leetcode

type TierNode struct {
	childs map[rune]*TierNode
	beLeaf bool
}

type TierTree struct {
	root *TierNode
}

func (t *TierTree) Insert(s string) {
	node := t.root
	for _, char := range s {
		if node.childs[char] == nil {
			node.childs[char] = &TierNode{
				childs: map[rune]*TierNode{},
			}
		}
		node = node.childs[char]
	}
	node.beLeaf = true
}

func searchTierTree(board [][]byte, row, col int, node *TierNode, pref string, result map[string]struct{}) {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] == '$' {
		return
	}
	char := rune(board[row][col])
	if node.childs[char] == nil {
		return
	}
	board[row][col] = '$'
	node = node.childs[char]
	pref += string(char)
	if node.beLeaf {
		result[pref] = struct{}{}
	}
	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, dir := range directions {
		searchTierTree(board, row+dir[0], col+dir[1], node, pref, result)
	}
	board[row][col] = byte(char)
}

func findWords(board [][]byte, words []string) []string {
	tt := &TierTree{
		root: &TierNode{
			childs: make(map[rune]*TierNode),
		},
	}
	for _, word := range words {
		tt.Insert(word)
	}
	m := len(board)
	n := len(board[0])
	result := map[string]struct{}{}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			searchTierTree(board, row, col, tt.root, "", result)
		}
	}
	ans := make([]string, 0, len(result))
	for s := range result {
		ans = append(ans, s)
	}
	return ans
}
