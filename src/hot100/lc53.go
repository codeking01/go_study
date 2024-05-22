package main

/**
 * @author xiongjl
 * @since 2024/1/3  10:34
 */

func maxSubArray(nums []int) int {
	// 直接动态规划
	// 方程 res[i]=max(res[i-1],nums[i)
	// 初始化 res[0]=nums[0]
	res := make([]int, len(nums))
	res[0] = nums[0]
	// 最大值
	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if res[i-1]+nums[i] > nums[i] {
			res[i] = res[i-1] + nums[i]
		} else {
			res[i] = nums[i]
		}
		if res[i] > maxValue {
			maxValue = res[i]
		}
	}
	return maxValue
}
