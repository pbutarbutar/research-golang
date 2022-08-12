package main

import "fmt"

func main() {
	fmt.Println(rotateMatrixShift1(3, 3, [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
	fmt.Println(rotateMatrixShift2([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
}

func rotateMatrixShift1(m int, n int, mat [][]int) (result [][]int) {
	row := 0
	col := 0
	var prev, curr int

	for row < m && col < n {
		if row+1 == m || col+1 == n {
			break
		}
		prev = mat[row+1][col]
		for i := col; i < n; i++ {
			curr = mat[row][i]
			mat[row][i] = prev
			prev = curr
		}
		row++
		for i := row; i < m; i++ {
			curr = mat[i][n-1]
			mat[i][n-1] = prev
			prev = curr
		}
		n--
		if row < m {
			for i := n - 1; i >= col; i-- {
				curr = mat[m-1][i]
				mat[m-1][i] = prev
				prev = curr
			}
		}
		m--
		if col < n {
			for i := m - 1; i >= row; i-- {
				curr = mat[i][col]
				mat[i][col] = prev
				prev = curr
			}
		}
		col++
	}
	result = mat
	return
}

func rotateMatrixShift2(matrix [][]int) [][]int {
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}
