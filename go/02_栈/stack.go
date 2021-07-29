package main

// 用数组实现一个基于字符串的固定大小的栈

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	ss := NewStringStack(3)

	for i := 0; i < 3; i++ {
		if !ss.Push("string") {
			log.Fatalln("入栈失败")
		}
	}

	for i := 0; i < 4; i++ {
		str, err := ss.Pull()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(str)
	}

}

type StringStacker interface {
	Push(string) bool
	Pull() (string, error)
}

type stringStack struct {
	items  []string
	count  int64 // 当前元素个数
	length int64 // 最大容量
}

func NewStringStack(n int64) StringStacker {
	return &stringStack{
		items:  make([]string, n),
		count:  0,
		length: n,
	}
}

func (s *stringStack) Push(item string) bool {

	if s.count == s.length {
		return false
	}

	s.items[s.count] = item
	s.count++

	return true
}

func (s *stringStack) Pull() (string, error) {

	if s.count == 0 {
		return "", errors.New("栈内无元素")
	}

	str := s.items[s.count-1]
	s.count--

	return str, nil

}
