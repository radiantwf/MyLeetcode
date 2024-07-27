package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("0-1"))
}

// 解题思路
// value*10 转化为位运算 (value<<3) + (value<<1)

func myAtoi(s string) int {
	var negative *bool
	var value int64
	for i := 0; i < len(s); i++ {
		if negative == nil && s[i] == ' ' {
			continue
		}
		if negative == nil && s[i] == '-' {
			n := true
			negative = &n
			continue
		}
		if negative == nil && s[i] == '+' {
			n := false
			negative = &n
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			break
		}
		if value > math.MaxInt32 {
			break
		}
		if negative == nil {
			n := false
			negative = &n
		}
		value = (value << 3) + (value << 1) + int64(s[i]-0x30)
	}
	if negative != nil && *negative {
		value = -value
	}
	if value < math.MinInt32 {
		value = math.MinInt32
	} else if value > math.MaxInt32 {
		value = math.MaxInt32
	}
	return int(value)
}
