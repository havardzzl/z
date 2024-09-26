package leetcode

// 1. 滑动窗口
// 找出无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[rune]bool)
	left, maxLen := 0, 0
	for right, char := range s {
		// 精髓
		for charSet[char] {
			delete(charSet, rune(s[left]))
			left++
		}
		charSet[char] = true
		if currentLen := right - left + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
