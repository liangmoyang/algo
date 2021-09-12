package main

import "fmt"

// 实现strstr()
// 力扣	https://leetcode-cn.com/problems/implement-strstr/

func main() {

	s1 := ""

	search := ""

	fmt.Println(strStr(s1, search))

}

func strStr(haystack string, needle string) int {

	// 父字符串不会是空字符串，或是空字符串，那么查找没有意义
	if needle == "" {
		return 0
	}

	m, n := len(haystack), len(needle)

outerFor:
	for i := 0; i+n <= m; i++ { // 假如遍历到i+n都已经超过了父串的长度，那就已经意味着父串中不可能包含子串
		for j := range needle {
			if haystack[i+j] != needle[j] { // 这里实现了在找到第一个字符之后，让父串可以跟子串同时递进
				continue outerFor // 直接跳转执行for的遍历
			}
		}
		return i
	}
	return -1
}
