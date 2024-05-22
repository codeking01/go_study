package main

/**
 * @author xiongjl
 * @since 2024/1/17  12:39
 */

// 这个可以用二分查找

func searchMatrix(matrix [][]int, target int) bool {
	// 先遍历行在遍历列
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] > target {
				break
			}
		}
	}
	/*for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[j][i] == target {
				return true
			} else if matrix[j][i] > target {
				break
			}
		}
	}*/
	return false
}
