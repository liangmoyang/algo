package main

import "fmt"

// 移动零，并保持原序
// 力扣：https://leetcode-cn.com/problems/move-zeroes/

func main() {

	arr := []int{9, 0, 0, 52, 7}

	moveZeroes(arr)

	fmt.Println(arr)

}

// 双指针解法在数组题中非常常见，需要理解两个指针的指向的目的
// 左指针指向的是【已经处理后的队列的尾部】
// 右指针指向的是【未处理的队列的头部】
func moveZeroes(nums []int) {

	left, right, length := 0, 0, len(nums) // 刚开始的时候，【已处理队列】为空，所以指针也指向第一位
	// 额外思考：如果前列多位非0，该如何减少对比次数

	for right < length {

		if nums[right] != 0 {

			// 右指针指向的不为0，则和左指针指向的当前元素换位
			nums[left], nums[right] = nums[right], nums[left]

			left++ // 左指针指向下一个尾部
		}
		right++
	}
}
