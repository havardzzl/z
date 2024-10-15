package leetcode

func addBinary(a string, b string) string {
	l1, l2 := len(a), len(b)
	lmin := min(l1, l2)
	lmax := max(l1, l2)
	res := make([]byte, lmax)
	ad := func(b1, b2 byte, plus int) (int, byte) {
		if b1 == '1' && b2 == '1' {
			if plus == 1 {
				return 1, '1'
			}
			return 1, '0'
		}
		if b1 == '0' && b2 == '0' {
			if plus == 1 {
				return 0, '1'
			}
			return 0, '0'
		}
		if plus == 1 {
			return 1, '0'
		}
		return 0, '1'
	}
	var i, plus int
	var rem byte
	for ; i < lmin; i++ {
		plus, rem = ad(a[l1-1-i], b[l2-1-i], plus)
		res[lmax-1-i] = rem
	}
	var remainBytes []byte
	if l1 > l2 {
		remainBytes = []byte(a)[:l1-lmin]
	} else {
		remainBytes = []byte(b)[:l2-lmin]
	}
	l := len(remainBytes)
	for i = 0; i < l; i++ {
		plus, rem = ad(remainBytes[l-1-i], '0', plus)
		res[l-1-i] = rem
	}
	if plus == 1 {
		return string(append([]byte{'1'}, res...))
	}
	return string(res)
}
