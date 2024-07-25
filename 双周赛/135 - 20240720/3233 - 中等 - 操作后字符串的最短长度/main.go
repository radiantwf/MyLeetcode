package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumLength("aa"))
}

// 解题思路：
// 当有奇数个相同字母时，只会保留1个
// 当有偶数个相同字母(>0)时，会保留2个
// 由于只包含小写字母，所以定义一个长度为26的计数器数组，统计每个字母出现的次数
// 遍历计数器数组，统计奇数次数的字母个数，偶数次数的字母个数*2

func minimumLength(s string) int {
	counterList := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counterList[s[i]-'a']++
	}
	counter := 0
	for i := 0; i < 26; i++ {
		if counterList[i]&1 == 1 {
			counter++
		} else if counterList[i] > 0 {
			counter += 2
		}
	}
	return counter
}
