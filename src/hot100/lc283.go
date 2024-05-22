package main

/**
 * @author xiongjl
 * @since 2023/12/26  14:26
 */

func moveZeroes(nums []int) {
	// 双指针
	l, r := 0, 0
	for _, item := range nums {
		if item == 0 {
			r++
		} else {
			nums[l] = nums[r]
			l++
			r++
		}
	}
	// 将l后面的全部置为0
	for l < r {
		nums[l] = 0
		l++
	}
}
