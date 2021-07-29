package main

import "fmt"

// 反转单链表
// 力扣：https://leetcode-cn.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode // 知识点：结构体自引用只能用指针
}

// 反转
func reverse(head *ListNode) *ListNode {

	// 声明一个前置节点。头部节点变为尾部节点需要指向一个nil指针
	var prev *ListNode

	// 将头部节点付给当前节点；
	// 每次循环都会把下一个节点复制给当前节点，以此遍历到尾部，当前节点指向原本的尾部节点的next
	curr := head

	for curr != nil {

		// 获取当前节点的下一个节点
		next := curr.Next

		// 反转：将当前节点的next指针指向上一个节点
		curr.Next = prev

		// 将当前节点变为前置节点
		prev = curr

		// 将下一个节点变为当前节点（相当于下次循环的当前节点）
		curr = next

	}

	// 反转到最后，最后一个前置指针，其实就是头部节点（原本的尾部节点）；curr指针在循环中已经指向了nil
	return prev
}

func main() {

	demo := []int{5, 4, 3, 2, 1}

	link := buildListNode(demo)

	newLink := reverse(link)

	printOneWayLinkedList(newLink)

}

// 打印单链表
func printOneWayLinkedList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val)
		list = list.Next
	}
}

// 将int切片构建成单链表
func buildListNode(list []int) *ListNode {
	if len(list) < 1 {
		return nil
	}

	res := &ListNode{
		Val: list[0],
	}

	temp := res

	// 注意这里用for range要跳过第一个，前面已经构建了一个头部节点，这里要从下标1开始
	for i := 1; i < len(list); i++ {
		temp.Next = &ListNode{
			Val: list[i],
		}
		temp = temp.Next
	}

	return res
}
