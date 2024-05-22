package main

/**
 * @author xiongjl
 * @since 2024/1/4  10:19
 */
// 超时，时间复杂度为o(kn)
/*func rotate(nums []int, k int) {
	// 需要取余数
	length := len(nums)
	k = k % length
	for i := 0; i < k; i++ {
		temp := nums[length-1]
		for j := length - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
}
*/

// 使用一个额外的数组
/*func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for index, item := range nums {
		newNums[(index+k)%len(nums)] = item
	}
	copy(nums, newNums)
}*/

// 翻转数组也可以
func _rotate(nums []int, k int) {

}
