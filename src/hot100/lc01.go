package main

/**
 * @author xiongjl
 * @since 2023/12/4  13:37
 */
func twoSum(nums []int, target int) []int {
	// a+b =target
	var numMap map[int]int
	numMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 存储key
		numMap[nums[i]] = i
	}
	// 判断 a=target-b
	for index, item := range nums {
		value, ok := numMap[target-item]
		if ok && value != index {
			return []int{index, value}
		}
	}
	return []int{}
}

/*func main() {
	twoSum([]int{3, 2, 4}, 6)
}
*/
