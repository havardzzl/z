package leetcode

import (
	"bytes"
	"sort"
)

func frequencySort(s string) string {
	record := map[byte]int{}
	for _, c := range []byte(s) {
		record[c]++
	}
	type pair struct {
		c byte
		q int
	}
	pairs := []pair{}
	for c, q := range record {
		pairs = append(pairs, pair{
			c: c,
			q: q,
		})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].q > pairs[j].q
	})
	res := []byte{}
	for _, p := range pairs {
		res = append(res, bytes.Repeat([]byte{p.c}, p.q)...)

	}
	return string(res)
}
