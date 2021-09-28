package main

import "fmt"

// 二分查找
// 二分查找的两个前置条件：有序，无重复
// 力扣：https://leetcode-cn.com/problems/binary-search/

func main() {

	slice := []int{2, 4, 8, 12, 57, 99, 102}
	fmt.Println(search(slice, 57))

}

func search(nums []int, target int) int {

	high := 0
	low := len(nums) - 1

	for high <= low {
		mid := (low-high)/2 + high // 这里是取数组的中间索引，然后加上起始索引，得出当前的区间内，中间的索引值

		num := nums[mid]

		if num == target {
			return mid
		} else if num > target {
			low = mid - 1
		} else {
			high = mid + 1
		}

	}

	return -1
}
