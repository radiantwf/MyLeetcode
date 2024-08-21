package main

import "fmt"

func init() {
}

func main() {
	fmt.Println(countKConstraintSubstrings("000011", 1))
}

func countKConstraintSubstrings(s string, k int) int {
	n := len(s)
	ret := n * (n + 1) / 2

	cnt0, cnt1 := 0, 0
	start := 0
	i := 0
	if s[0] == '0' {
		cnt0++
	} else {
		cnt1++
	}
	for {
		if cnt0 > k && cnt1 > k {
			ret -= (n - i)
			if s[start] == '0' {
				cnt0--
			} else {
				cnt1--
			}
			start++
		} else {
			i++
			if i == n {
				break
			}
			if s[i] == '0' {
				cnt0++
			} else {
				cnt1++
			}
		}
	}

	return ret
}
