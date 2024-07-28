package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nonSpecialCount(1, 2))
}

// 解题思路
// 质数只有一个真因数1；合数有多个真因数，只包含2个真因数的合数只有可能是>1质数的平方数。

var primesCounter []int

func init() {
	max := int(math.Sqrt(math.Pow10(9)))
	primesCounter = make([]int, max+1)
	for i := 2; i <= max; i++ {
		if primesCounter[i] == 0 {
			primesCounter[i] = primesCounter[i-1] + 1
			for j := i * i; j <= max; j += i {
				primesCounter[j] = -1
			}
		} else {
			primesCounter[i] = primesCounter[i-1]
		}
	}
}

func nonSpecialCount(l int, r int) int {
	counterR := primesCounter[int(math.Sqrt(float64(r)))]
	countL := primesCounter[int(math.Sqrt(float64(l-1)))]
	return r - l + 1 - (counterR - countL)
}

// map保存平方数，会导致超时
// var primesPow2 map[int]bool

// func init() {
// 	max := int(math.Sqrt(math.Pow10(9)))
// 	isPrime := make([]bool, max+1)
// 	isPrime[2] = true
// 	for i := 3; i <= max; i += 2 {
// 		isPrime[i] = true
// 	}
// 	for i := 3; i <= max; i += 2 {
// 		if isPrime[i] {
// 			for j := i * i; j <= max; j += i {
// 				isPrime[j] = false
// 			}
// 		}
// 	}
// 	primesPow2 = make(map[int]bool)
// 	primesPow2[2*2] = true
// 	for i := 3; i <= max; i += 2 {
// 		if isPrime[i] {
// 			primesPow2[i*i] = true
// 		}
// 	}
// }

// func nonSpecialCount(l int, r int) int {
// 	count := 0
// 	for i := l; i <= r; i++ {
// 		if _, ok := primesPow2[i]; ok {
// 			count++
// 		}
// 	}
// 	return r - l + 1 - count
// }
