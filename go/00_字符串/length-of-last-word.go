package main

import (
	"fmt"
	"strings"
)

// 最后一个单词的长度
// 力扣 https://leetcode-cn.com/problems/length-of-last-word/

// 题目很简单，但是解法有很多种
func main() {

	//str := "   fly me   to   the moon  "
	str := "a"
	fmt.Println(lengthOfLastWord_1(str))
	fmt.Println(lengthOfLastWord_2(str))
}

// 最简单解法，用库的优势
func lengthOfLastWord_1(s string) int {
	arr := strings.Split(strings.TrimSpace(s), " ")

	return len(arr[len(arr)-1])

}

// 循环解法
func lengthOfLastWord_2(s string) int {

	var length int

	// 从尾部开始循环
	for i := len(s) - 1; i >= 0; i-- {

		if s[i] != ' ' {
			length++
			continue
		}

		if s[i] == ' ' && length > 0 {
			break
		}

	}

	return length

}
