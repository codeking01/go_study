package main

/**
 * @author xiongjl
 * @since 2024/1/6  17:38
 */

func setZeroes(matrix [][]int) {
	xLen, yLen := len(matrix), len(matrix[0])
	xzero, yzero := matrix[0][0], matrix[0][0]
	// 遍历第一行，第一列
	for i := 0; i < yLen; i++ {
		if matrix[0][i] == 0 {
			xzero = 0
			break
		}
	}
	for i := 0; i < xLen; i++ {
		if matrix[i][0] == 0 {
			yzero = 0
			break
		}
	}
	// 上边记录0是当前列需要置为0，左边一列是当前这一行需要置为0
	for i := 1; i < xLen; i++ {
		for j := 1; j < yLen; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for j := 1; j < yLen; j++ {
		if matrix[0][j] == 0 {
			// 这一列置0
			for k := 0; k < xLen; k++ {
				matrix[k][j] = 0
			}
		}
	}
	for j := 1; j < xLen; j++ {
		if matrix[j][0] == 0 {
			// 置0
			for k := 0; k < yLen; k++ {
				matrix[j][k] = 0
			}
		}
	}
	if xzero == 0 {
		for i := 0; i < yLen; i++ {
			matrix[0][i] = 0
		}
	}
	if yzero == 0 {
		for i := 0; i < xLen; i++ {
			matrix[i][0] = 0
		}
	}
}
