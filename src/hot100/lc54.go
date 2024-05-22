package main

/**
 * @author xiongjl
 * @since 2024/1/14  19:46
 */

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	top, right, down, left := 0, n-1, m-1, 0
	res := make([]int, 0)
	for {
		// 上右
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		if top > down {
			break
		}
		// 右下
		for i := top; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// 下左
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if down < top {
			break
		}
		// 左上
		for i := down; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
