package main

// 整数反转
// 力扣：https://leetcode-cn.com/problems/reverse-integer/

import (
	"fmt"
	"math"
)

func main() {
	i := -99591
	fmt.Println(reverse(i))
}

// 易于理解版
func reverse(x int) int {

	if x == 0 {
		return 0
	}

	res := 0 // res是正式返回的值

	for {

		// 假设此时res再次*/10就超出了32位整数的范围，那么就没必要再走下去了
		// 因为后续res会执行*10，所以哪怕r最后是0，也会溢出
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}

		r := x % 10 // 每次只需要取出最后一位数字
		x /= 10     // 取出最后一位数字后，就可以把基数除以10

		if r == x { //
			break
		}

		res = res*10 + r

	}

	return res
}
