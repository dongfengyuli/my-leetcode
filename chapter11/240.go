package main

import "fmt"

func main(){
	matrix := [][]int{
		{1,4,7,11,15},
		{2,5,8,12,19},
		{3,6,9,16,22},
		{10,13,14,17,24},
		{18,21,23,26,30},
	}

	res := searchMatrix(matrix,5)
	fmt.Println(res)
}


func searchMatrix(matrix [][]int, target int) bool{
	if len(matrix) == 0{
		return false
	}

	col,row := 0,len(matrix) - 1
	for row >= 0 && col < len(matrix[row]){
		if matrix[row][col] == target{
			return true
		}else if matrix[row][col] < target{
			col++
		}else{
			row--
		}
	}
	return false
}