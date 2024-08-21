package main

import "fmt"

func init() {
}

func main() {
	fmt.Println(countKConstraintSubstrings("000011", 1))
}

func countKConstraintSubstrings(s string, k int) int {
	// 计算前缀和
	n := len(s)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + int(s[i-1]-'0')
	}

	ret := 0
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			sum := sum[j] - sum[i-1]
			if sum <= k || j-i+1-sum <= k {
				ret++
			} else {
				break
			}
		}
	}

	return ret
}
