package main

import (
	"fmt"
	"math"
)

func main() {
	// 乘方
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(-2147483648, -1))
}

// 解题思路
// 由于数值范围是 [−2^31,  2^31−1]，所以直接使用int64类型进行计算
// 被除数、除数都转换为正数进行计算，最后再设置符号
// 由于不能使用乘法，所以使用位运算进行乘法计算
// 被除数、除数 同时向左移位，直到被除数小于除数
// 每次移位时，记录移位次数，同时记录移位的值
// 被除数减去移位的值，继续移位
// 重复上述步骤，直到被除数小于除数
// 每次循环 1<<位移次数 相加，即为商
// 设置商的符号
// 如果结果值超出int32范围，则返回int32的最大值或最小值，否则返回结果值

func divide(dividend int, divisor int) int {
	negative := false
	v1 := int64(dividend)
	v2 := int64(divisor)
	if v1 < 0 {
		v1 = -v1
		negative = !negative
	}
	if v2 < 0 {
		v2 = -v2
		negative = !negative
	}
	if v1 < v2 {
		return 0
	}
	value := int64(0)
	for v1 >= v2 {
		temp := v2
		n := 0
		for v1 >= (temp << 1) {
			temp <<= 1
			n++
		}
		v1 -= temp
		value += 1 << n
	}

	if negative {
		value = -value
	}
	if value < math.MinInt32 {
		value = math.MinInt32
	} else if value > math.MaxInt32 {
		value = math.MaxInt32
	}

	return int(value)
}
