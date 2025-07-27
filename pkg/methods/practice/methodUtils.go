package main

import "fmt"

type MethodUtils struct {
	matrix [3][3]int
}

func (mu MethodUtils) CallMethod(start int, end int) {
	if start > end {
		fmt.Printf("start %d is greater than end %d\n", start, end)
	}
	for i := start; i <= end; i++ {
		for j := start; j <= i; j++ {
			fmt.Printf("%d✖️%d=%d\t", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}

func (mu *MethodUtils) transpose() [][]int {
	matrix := mu.matrix
	transposeMatrix := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		row := make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			row[j] = matrix[j][i]
		}
		transposeMatrix[i] = row
	}
	return transposeMatrix
}

func main() {
	mu := MethodUtils{
		matrix: [3][3]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}
	mu.CallMethod(1, 9)
	mu.CallMethod(3, 8)
	res := mu.transpose()
	fmt.Printf("result is : %v", res)
}
