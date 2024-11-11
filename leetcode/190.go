package leetcode

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		if num%2 == 1 {
			ans += 1 << (31 - i)
		}
		num = num >> 1
	}
	return ans
}
