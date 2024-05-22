package main

import (
	"sort"
)

/**
 * @author xiongjl
 * @since 2023/12/26  20:21
 */

func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	var result [][]int
	length := len(nums)
	for i := 0; i < length; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] || a > 0 {
			continue
		}
		for l, r := i+1, length-1; l < r; {
			if a+nums[l]+nums[r] > 0 {
				r--
			} else if a+nums[l]+nums[r] < 0 {
				l++
			} else {
				temp := append([]int{}, a, nums[l], nums[r])
				result = append(result, temp)
				// 去重
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
				r--
			}
		}
	}
	return result
}

/*func main() {
	test := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(test)
	fmt.Println(result)
}
*/
