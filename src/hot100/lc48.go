package main

/**
 * @author xiongjl
 * @since 2024/1/15  15:34
 */

func rotate(matrix [][]int) {
	//n*n的矩阵
	n := len(matrix) - 1
	// 行
	for i := 0; i < (n+1)/2; i++ {
		// 列
		for j := 0; j <= n/2; j++ {
			// 记录左上角的
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = matrix[j][n-i]
			matrix[j][n-i] = temp
		}
	}
}
