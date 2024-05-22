package main

/**
 * @author xiongjl
 * @since 2024/1/5  10:43
 */

/*
	func productExceptSelf(nums []int) []int {
		// 两个数组，一个左边的乘积，一个右边的城际
		leftNums, rightNums, res := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
		// 初始化
		leftNums[0], rightNums[len(rightNums)-1] = 1, 1
		// 左数组
		for i := 1; i < len(leftNums); i++ {
			leftNums[i] = nums[i-1] * leftNums[i-1]
		}
		// 右数组
		for i := len(rightNums) - 2; i >= 0; i-- {
			rightNums[i] = nums[i+1] * rightNums[i+1]
		}
		// 结果
		for i := 0; i < len(leftNums); i++ {
			res[i] = leftNums[i] * rightNums[i]
		}
		return res
	}
*/
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	// 先存储左乘积
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	// 右边的乘积依次计算
	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * r
		r *= nums[i]
	}
	return res
}
