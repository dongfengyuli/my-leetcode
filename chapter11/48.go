package main

import "fmt"

func main(){

	matrix := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	rotate(matrix)

	fmt.Println(matrix)

}


func rotate(matrix [][]int)  {
	length := len(matrix)

	// 先左<->右互换
	for i := 0; i < length; i++ {
		for j := 0; j < length / 2; j++ {
			matrix[i][j], matrix[i][length - 1 - j] = matrix[i][length - 1 - j], matrix[i][j]
		}
	}

	// 后关于左下/右上的对角线互换
	for i := 0; i < length - 1; i++ {
		for j := 0; j < length - i - 1; j++ {
			matrix[i][j], matrix[length - j - 1][length - i - 1] = matrix[length - j - 1][length - i - 1], matrix[i][j]
		}
	}

}