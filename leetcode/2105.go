package leetcode

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	l, r := 0, len(plants)-1
	capL, capR := capacityA, capacityB
	ans := 0
	for l < r {
		if capL < plants[l] {
			capL = capacityA
			ans++
		}
		capL -= plants[l]
		if capR < plants[r] {
			capR = capacityB
			ans++
		}
		capR -= plants[r]
		l++
		r--
	}
	if l == r && capL < plants[l] && capR < plants[l] {
		ans++
	}
	return ans
}
