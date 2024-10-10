package leetcode

import "strings"

func fullFormat(words []string, normalWidth, maxWidth int) string {
	extraBlank := maxWidth - normalWidth
	n := len(words)
	if n == 1 {
		return words[0] + strings.Repeat(" ", extraBlank)
	}
	sb := strings.Builder{}
	sb.WriteString(words[0])
	sep := strings.Repeat(" ", 1+extraBlank/(n-1))
	for i := 0; i < extraBlank%(n-1); i++ {
		sb.WriteString(sep + " " + words[i+1])
	}
	for i := extraBlank % (n - 1); i < n-1; i++ {
		sb.WriteString(sep + words[i+1])
	}
	return sb.String()
}

func fullJustify(words []string, maxWidth int) []string {
	ans := []string{}
	parts := []string{}
	n := len(words)

	for idx := 0; idx < n; idx++ {
		parts = append(parts, words[idx])
		l := len(words[idx])
		for l <= maxWidth && idx+1 < n {
			idx++
			parts = append(parts, words[idx])
			l += len(words[idx]) + 1
		}
		if l > maxWidth {
			parts = parts[:len(parts)-1]
			l -= len(words[idx]) + 1
			idx--
			ans = append(ans, fullFormat(parts, l, maxWidth))
		} else {
			ans = append(ans, strings.Join(parts, " ")+strings.Repeat(" ", maxWidth-l))
		}
		parts = parts[:0]
	}
	return ans
}
