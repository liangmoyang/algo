package main

// 有效的括号
// 力扣：https://leetcode-cn.com/problems/valid-parentheses/
// 对于栈结构来说是比较基础的题目，难点在于

import "fmt"

func main() {

	//str := "()[]{}"
	str := "([)]"

	fmt.Println(isValid(str))

}

func isValid(s string) bool {

	n := len(s)

	// 奇数的字符串肯定不能完全匹配
	if n%2 == 1 {
		return false
	}

	// 使用map存储，可以便捷查到找对应的符号，免去多个判断
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 定义一个变量，作为栈
	stack := []byte{}

	for i := 0; i < n; i++ {

		// 这里是配合map筛选了前置，只有后置符号能进来
		if pairs[s[i]] > 0 {

			// 条件一，后置的进来了，然而栈内并没有前置，那就说明字符串并不符合，所以false
			// 条件二，用栈内的最后一个（即栈顶元素），和当前后置符号的前置（从map中用后置取前置）做对比，如果不一致，则说明不能抵消
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}

			// 经过前面的筛选，说明当前字符跟栈顶元素能匹配，则把栈顶元素出栈
			stack = stack[:len(stack)-1]
		} else {
			// 当前字符是前置符号，直接压栈
			stack = append(stack, s[i])
		}
	}

	// 执行完了之后，如果栈内还有前置元素，则说明不匹配
	return len(stack) == 0

}
