package main

import "fmt"

// 移除元素
// 力扣：https://leetcode-cn.com/problems/remove-element/

func main() {
	slice := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//fmt.Println(removeElement1(slice,2))
	//fmt.Println(removeElement2(slice,2))
	fmt.Println(removeElement3(slice, 2))
	fmt.Println(slice)
}

// 暴力破解法 空间O(1) 时间O(n^2)
func removeElement1(nums []int, val int) int {

	count := len(nums)

	for i := 0; i < count; i++ {

		if nums[i] == val {

			for j := i + 1; j < count; j++ {
				nums[j-1] = nums[j]
			}

			i-- // 因为都往前迁了一位，这里所以这里要减去一位，才是原本的下一位，不然会跳过一个元素
			count--
		}
	}

	return count
}

// 双指针解法1 空间O(1) 空间O(n)
func removeElement2(nums []int, val int) int {

	left := 0

	for _, v := range nums { // 利用range的遍历，这里的v也相当于是右指针了
		if v != val {
			nums[left] = v // 把所有不等于val的数都逐个移到前面，等于val的元素会被下一个v替换掉；这也会导致前面的元素会重复赋值
			left++         // 左指针右移，左边的都不包含val
		}
	}

	return left
}

// 双指针解法2 空间O(1) 空间O(n)
func removeElement3(nums []int, val int) int {
	left, right := 0, len(nums)

	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
