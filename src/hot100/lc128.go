package main

import (
	"fmt"
)

/**
 * @author xiongjl
 * @since 2023/12/26  13:15
 */

func longestConsecutive(nums []int) int {
	// 直接放字典里面依次遍历就行
	m := make(map[int]int)
	for _, num := range nums {
		m[num] = 1
	}
	max := 0
	for key := range m {
		// 找边界值，其他的不用判断
		if m[key-1] == 0 {
			fmt.Println("m[key-1]", m[key-1])
			cnt := 1
			// 最小边界值
			for m[key+1] == 1 {
				key++
				cnt++
			}
			if cnt > max {
				max = cnt
			}
		}
	}
	return max
}

//func main() {
//	nums := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
//	res := longestConsecutive(nums)
//	fmt.Println(res)
//}
