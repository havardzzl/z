package leetcode

import "strings"

func simplifyPath(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	simplifyParts := []string{}
	for _, part := range parts {
		switch part {
		case ".", "":
		case "..":
			if len(simplifyParts) > 0 {
				simplifyParts = simplifyParts[:len(simplifyParts)-1]
			}
		default:
			simplifyParts = append(simplifyParts, part)
		}
	}
	return "/" + strings.Join(simplifyParts, "/")
}
