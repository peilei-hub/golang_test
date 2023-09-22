package main

import "fmt"

// https://leetcode.cn/problems/powx-n/

// x^10 = x^5 * x^5
// x^5 = x^2 * x^2 * x
// n小于0，则结果求倒数

func pow50(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	v := pow50(x, n/2)

	if n%2 != 0 { // 是奇数
		return v * v * x
	}
	return v * v // 是偶数
}

func myPow(x float64, n int) float64 {
	absN := n
	if n < 0 {
		absN = -n
	}

	res := pow50(x, absN)

	if n < 0 {
		return 1 / res
	}
	return res
}

func main() {
	pow := myPow(2, 3)
	fmt.Println(pow)
}
