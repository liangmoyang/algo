package main

// 回文数
// 力扣：https://leetcode-cn.com/problems/palindrome-number/submissions/

import "fmt"

func main() {

	x := 12321

	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {

	// 按照题解，负数都不是回文数
	if x < 0 {
		return false
	}

	// 尾数为0的，也不是回文数
	if x%10 == 0 && x != 0 { // 注意这里要筛掉0，因为0是回文数
		return false
	}

	newNumber := 0
	// 接着来反转源数字的一半
	for x > newNumber {

		lastNum := x % 10

		x = x / 10

		newNumber = newNumber*10 + lastNum

	}

	// 前部分是长度为偶数的情况：如果源数字的一半等于反转后的数字成立
	// 后部分是长度为奇数的情况，所以要把反转后的数字去掉最后一个数字
	return x == newNumber || x == newNumber/10
}
