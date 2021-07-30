package main

// 求两数之和
// 力扣：https://leetcode-cn.com/problems/two-sum/

import "fmt"

func main() {

	testNums := []int{2, 5, 7, 6, 8}

	res := towSumGood(testNums, 9)

	fmt.Println(res)
}

// twoSum 容易理解的入门版本
func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))
	for k, v := range nums {
		m[v] = k
	}

	for k, v := range nums {
		if i, ok := m[target-v]; ok && k < i {
			return []int{k, i}
		}
	}

	return nil
}

// twoSumGood 一个好的版本
func towSumGood(nums []int, target int) []int {

	m := make(map[int]int)

	for k, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{index, k}
		}

		m[v] = k
	}

	return nil
}
