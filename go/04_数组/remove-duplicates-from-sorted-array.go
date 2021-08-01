package main

// 删除有序数组中的重复项
// 力扣：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 力扣的题目描述不太清晰，按题解回推，题目只是要求“元素不重复”，而不是要求“把原数组的重复值删除”
// 所以在题解里是把处理过的元素值前置，数组的原长度并没有改变

import "fmt"

func main() {
	exmp := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4}

	fmt.Println(removeDuplicates(exmp))

}

// 对于“有序”的数组，很多情况下用双指针都可以解决
// 这里双指针的解法是逐个对比
// 如果后方跟前方不一致，则把后方赋值给前方
// 然后左指针右移，直到找到下一个和前方不一样的值，再赋值给左指针的值
// 这样左边的值就只会出现一次，从而让左指针指向未重复数组的尾部，也刚好是长度
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	left, right := 1, 1
	for right < n {

		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left // 这里返回左指针而不是长度的原因是因为在遍历过后，左指针指向的就是处理过后的数组的尾部
}
