package main

import "fmt"

var a1 [][]int

func main() {
	fmt.Println("Base Array Matrix: ", [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	})

	fmt.Println("Shift 1: ", rotateMatrixShift1(3, 3, [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}))

	fmt.Println("Shift 2: ", rotateMatrixShift2([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}))
}

func rotateMatrixShift1(m int, n int, mat [][]int) [][]int {
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

	return mat
}

func rotateMatrixShift2(mat [][]int) [][]int {
	n := len(mat)
	small := n / 2
	large := small
	if n%2 != 0 {
		large = small + 1
	}
	for i := 0; i < large; i++ {
		for j := 0; j < small; j++ {
			temp := mat[i][j]
			mat[i][j] = mat[n-1-j][i]
			mat[n-1-j][i] = mat[n-1-i][n-1-j]
			mat[n-1-i][n-1-j] = mat[j][n-1-i]
			mat[j][n-1-i] = temp
		}
	}
	return mat
}
