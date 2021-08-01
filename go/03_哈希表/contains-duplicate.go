package main

import "fmt"

// 存在重复元素（这个是我认为在力扣上最简单的题目，不需要其他的想法，只要理解了哈希的特性就能做出来）
// 力扣：https://leetcode-cn.com/problems/contains-duplicate/

func main() {
	exmp := []int{1, 8, 7, 7, 2, 4, 6, 2, 0, 5, 7}
	fmt.Println(containsDuplicate(exmp))
}

func containsDuplicate(nums []int) bool {

	m := make(map[int]struct{}) // 空结构体不占用空间

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}

		m[v] = struct{}{}
	}

	return false
}
