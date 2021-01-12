package Day116

import "fmt"

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	if row == 0 {
		return NumMatrix{}
	}
	col := len(matrix[0])
	if col == 0 {
		return NumMatrix{}
	}
	m := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		m[i] = make([]int, col+1)
		if i == 0 {
			continue
		}
		for j := 1; j <= col; j++ {
			m[i][j] = m[i][j-1] + matrix[i-1][j-1]
		}
		fmt.Printf("%+v\n", m[i])
	}
	for j := 1; j <= col; j++ {
		for i := 1; i <= row; i++ {
			m[i][j] += m[i-1][j]
		}
	}
	fmt.Printf("======\n")
	for i := 0; i <= row; i++ {
		fmt.Printf("%+v\n", m[i])
	}
	return NumMatrix{
		matrix: m,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if len(this.matrix) == 0 {
		return 0
	}
	return this.matrix[row2+1][col2+1] + this.matrix[row1][col1] - this.matrix[row1][col2+1] - this.matrix[row2+1][col1]
}
