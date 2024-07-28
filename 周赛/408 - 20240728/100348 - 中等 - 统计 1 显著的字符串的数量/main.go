package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nonSpecialCount(4, 16))
}

// 解题思路
// 质数只有一个真因数1；合数有多个真因数，只包含2个真因数的合数只有可能是>1质数的平方数。

func nonSpecialCount(l int, r int) int {
	limit := int(math.Sqrt(float64(r)))
	begin := int(math.Sqrt(float64(l)))

	primes := getPrimes(limit)
	count := 0
	for i := begin; i <= limit; i++ {
		if _, ok := primes[i]; ok {
			if i == begin {
				if math.Pow(float64(i), 2) == float64(l) {
					count++
				} else {
					continue
				}
			}
			count++
		}
	}

	return count
}

func getPrimes(max int) map[int]bool {
	isPrime := make([]bool, max+1)
	isPrime[2] = true
	for i := 3; i <= max; i += 2 {
		isPrime[i] = true
	}
	for i := 3; i <= max; i += 2 {
		if isPrime[i] {
			for j := i * i; j <= max; j += i {
				isPrime[j] = false
			}
		}
	}
	primes := make(map[int]bool)
	primes[2] = true
	for i := 3; i <= max; i += 2 {
		if isPrime[i] {
			primes[i] = true
		}
	}
	return primes
}
