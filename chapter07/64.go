package main

import "fmt"

func main(){
	grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}

	res := minPathSum(grid)

	fmt.Println(res)
}



func minPathSum(grid [][]int) int{
	length := len(grid)
	if length < 1{
		return 0
	}

	for i:=0;i<length;i++{
		for j :=0;j<len(grid[i]);j++{
			if i == 0 && j!=0{
				grid[i][j] = grid[i][j-1] + grid[i][j]
			}else if j == 0 && i != 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			}else if i !=0 && j !=0 {
				grid[i][j] = min(grid[i-1][j],grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[length-1][len(grid[length - 1]) - 1]
}

func min(a, b int)  int{
	if a > b{
		return b
	}
	return a
}