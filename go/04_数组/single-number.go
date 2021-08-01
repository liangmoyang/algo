package main

import "fmt"

// 只出现一次的数字
// 力扣：https://leetcode-cn.com/problems/single-number/

func main() {

	//exmp := []int{2,2,1}
	exmp := []int{8, 5, 8, 1, 5, 7, 1}

	num := singleNumber(exmp)

	fmt.Println(num)
	testXOR()
}

func singleNumber(nums []int) int {

	res := 0

	for _, v := range nums {
		res = res ^ v // 异或的精髓：异或同一个数等于没有异或（所以在原题中，必须是其余元素出现双数次）
	}

	return res
}

func testXOR() {
	fmt.Println(6 ^ 6 ^ 7) // 7
}
