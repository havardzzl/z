package leetcode

import "strings"

func decodeString(s string) string {
	var cnt int
	type node struct {
		cnt  int
		part []byte
	}
	st := []node{}
	sb := strings.Builder{}
	for _, c := range []byte(s) {
		switch {
		case c == '[':
			st = append(st, node{cnt: cnt, part: []byte{}})
			cnt = 0
		case c == ']':
			sti := st[len(st)-1]
			parts := []byte{}
			for i := 0; i < sti.cnt; i++ {
				parts = append(parts, sti.part...)
			}
			st = st[:len(st)-1]
			if len(st) == 0 {
				sb.Write(parts)
			} else {
				st[len(st)-1].part = append(st[len(st)-1].part, parts...)
			}
		case c >= '0' && c <= '9':
			cnt = cnt*10 + int(c-'0')
		default:
			if len(st) == 0 {
				sb.WriteByte(c)
			} else {
				st[len(st)-1].part = append(st[len(st)-1].part, c)
			}
		}
	}
	return sb.String()
}
