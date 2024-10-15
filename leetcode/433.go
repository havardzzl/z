package leetcode

import "bytes"

func minMutation(startGene string, endGene string, bank []string) int {
	banks := make(map[string]bool, len(bank))
	for _, b := range bank {
		banks[b] = true
	}
	visit := map[string]bool{}
	var step int
	gs := []string{startGene}
	bs := []byte{'A', 'C', 'G', 'T'}
	fwd := func(gs []string) []string {
		newGs := []string{}
		for _, gene := range gs {
			genes := []byte(gene)
			for i, b := range genes {
				idx := bytes.IndexByte(bs, b)
				for j := 1; j < 4; j++ {
					mut := bytes.Clone(genes)
					mut[i] = bs[(idx+j)%4]
					smut := string(mut)
					if banks[smut] && !visit[smut] {
						delete(banks, smut)
						visit[smut] = true
						newGs = append(newGs, smut)
					}
				}
			}
		}
		return newGs
	}
	for len(gs) > 0 {
		gs = fwd(gs)
		step++
		if visit[endGene] {
			return step
		}
	}
	return -1
}
