package leetcode

import "fmt"

func isPalindrome(s string) bool {
	bs := []byte(s)
	bsClean := []byte{}
	for _, b := range bs {
		if b <= 'Z' && b >= 'A' {
			bsClean = append(bsClean, b-'A'+'a')
		} else if b <= '9' && b >= '0' {
			bsClean = append(bsClean, b)
		} else if b <= 'z' && b >= 'a' {
			bsClean = append(bsClean, b)
		}
	}
	fmt.Println(string(bsClean))
	for i := 0; i < len(bsClean)/2; i++ {
		if bsClean[i] != bsClean[len(bsClean)-1-i] {
			return false
		}
	}
	return true
}
