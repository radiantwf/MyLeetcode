package main

import "fmt"

func main() {
	s := "aa"
	p := "a*"
	ret := isMatch(s, p)
	fmt.Println(ret)
}

func isMatch(s string, p string) bool {
	sIdxEnd, pIdxEnd := len(s), len(p)
	for sIdxEnd > 0 && pIdxEnd > 0 {
		if p[pIdxEnd-1] == '*' {
			break
		}
		if s[sIdxEnd-1] != p[pIdxEnd-1] && p[pIdxEnd-1] != '.' {
			return false
		}
		sIdxEnd--
		pIdxEnd--
	}
	if sIdxEnd == 0 && pIdxEnd == 0 {
		return true
	} else if pIdxEnd > 0 && p[pIdxEnd-1] != '*' {
		return false
	}

	return false
}

func asteriskMatch(s, p string, inputSIdxEnd, inputPIdxEnd int) (sIdxEnd, pIdxEnd int, ret bool) {
	return
}
