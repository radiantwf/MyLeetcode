package main

import (
	"fmt"
	"strings"
)

func init() {
}

func main() {
	fmt.Println(largestPalindrome(1, 2))
}

func largestPalindrome(n int, k int) string {
	ret := strings.Builder{}
	switch k {
	// 1, 3, 9 结果是每位都为9
	case 1:
		fallthrough
	case 3:
		fallthrough
	case 9:
		ret.WriteString(strings.Repeat("9", n))
	// 4 判断最后两位数字被4整除的最大回文数字 为88
	// 所以大于4位的数字，中间部分为9，前后部分为88
	// 1位数字为8，2位数字为88，3位数字为888，4位数字为8888
	case 4:
		if n <= 4 {
			ret.WriteString(strings.Repeat("8", n))
		} else {
			ret.WriteString("88")
			ret.WriteString(strings.Repeat("9", n-4))
			ret.WriteString("88")
		}
	// 8 判断最后三位数字被8整除的最大回文数字 为888
	// 所以大于6位的数字，中间部分为9，前后部分为888
	// 1位数字为8，2位数字为88，3位数字为888，4位数字为8888，5位数字为88888，6位数字为888888
	case 8:
		if n <= 6 {
			ret.WriteString(strings.Repeat("8", n))
		} else {
			ret.WriteString("888")
			ret.WriteString(strings.Repeat("9", n-6))
			ret.WriteString("888")
		}
	// 2 判断最后一位数字被2整除的最大回文数字 为8
	// 所以位数大于2位的数字，中间部分为9，前后部分为8
	// 1位数字为8，2位数字为88
	case 2:
		fallthrough
	// 5 判断最后一位数字被2整除的最大回文数字 为5
	// 所以位数大于2位的数字，中间部分为9，前后部分为5
	// 1位数字为5，2位数字为55
	case 5:
		if k == 2 {
			k = 8
		}
		ret.WriteByte(byte(k + 0x30))
		if n > 2 {
			ret.WriteString(strings.Repeat("9", n-2))
		}
		if n > 1 {
			ret.WriteByte(byte(k + 0x30))
		}
	// 6 判断数字同时被2、3整除
	// 被2整除，判断最后一位数字是2的倍数即可
	// 被3整除，判断所有数字之和是3的倍数即可
	// 所以位数大于2位的数字，前后部分为8 可被2整除的最大数字，奇数位数，中间1位填充8，偶数位数，中间两位填充77，其余部分填充9
	// 1位数字为6，2位数字为66
	case 6:
		if n <= 2 {
			ret.WriteString(strings.Repeat("6", n))
		} else {
			suf := strings.Builder{}
			ret.WriteString("8")
			ret.WriteString(strings.Repeat("9", (n-1)/2-1))
			suf.WriteString(strings.Repeat("9", (n-1)/2-1))
			if n%2 == 0 {
				ret.WriteString("77")
			} else {
				ret.WriteString("8")
			}
			ret.WriteString(suf.String())
			ret.WriteString("8")
		}
	// 判断数字被7整除，没发现具体规律，只能部分暴力解决
	// 可以被7整除大于2位的数字，两端为9，只修改中间部分数据（1位或2位）即可满足要求，所以只需从两端开始设置9，求中间部分的最大数字即可
	// 1位数字为7，2位数字为77
	case 7:
		if n <= 2 {
			ret.WriteString(strings.Repeat("7", n))
		} else {
			var checkFunc func(string, int, int) string
			checkFunc = func(num string, pre, index int) string {
				for max := 9; max >= 0; max-- {
					new_pre := (pre*10 + max) % 7
					new_num := num + string(max+0x30)
					if index == (n-1)/2 {
						var startIndex int
						if n%2 == 0 {
							startIndex = len(new_num) - 1
						} else {
							startIndex = len(new_num) - 2
						}
						for i := startIndex; i >= 0; i-- {
							new_pre = (new_pre*10 + int(new_num[i]-0x30)) % 7
						}
						if new_pre == 0 {
							numBuilder := strings.Builder{}
							numBuilder.WriteString(new_num)
							for i := startIndex; i >= 0; i-- {
								numBuilder.WriteByte(new_num[i])
							}
							return numBuilder.String()
						}
					} else {
						ret := checkFunc(new_num, new_pre, index+1)
						if ret != "" {
							return ret
						}
					}
				}
				return ""
			}
			pre := 0
			for i := 0; i < (n-1)/2; i++ {
				pre = (pre*10 + 9) % 7
			}
			return checkFunc(strings.Repeat("9", (n-1)/2), pre, (n-1)/2)
		}
	}

	return ret.String()
}
