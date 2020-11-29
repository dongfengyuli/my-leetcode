package main

import "fmt"

func main(){
	matrix := [][]int{
		{1,2,2,3,5},
		{3,2,3,4,4},
		{2,4,5,3,1},
		{6,7,1,4,5},
		{5,1,1,2,4},
	}
	res := pacificAtlantic(matrix)
	fmt.Println(res)
}


func pacificAtlantic(matrix [][]int) [][]int {

	res := make([][]int, 0)
	if nil == matrix || len(matrix) == 0 {
		return res
	}

	row, col, dirs := len(matrix), len(matrix[0]), [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	inArea := func(x, y int) bool {
		return x >= 0 && x < row && y >= 0 && y < col
	}
	var dfs func(x, y int, matrix [][]int, travelled [][]int)
	dfs = func(x, y int, matrix [][]int, travelled [][]int) {
		if travelled[x][y] == 1 {
			return
		}
		travelled[x][y] = 1
		for _, d := range dirs {
			newx := x + d[0]
			newy := y + d[1]
			if !inArea(newx, newy) || matrix[x][y] > matrix[newx][newy] || travelled[newx][newy] == 1 {
				continue
			}
			dfs(newx, newy, matrix, travelled)
		}
	}


	pcf, atl := make([][]int, row), make([][]int, row)
	for i := 0; i < row; i++ {
		pcf[i] = make([]int, col)
		atl[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		dfs(i, 0, matrix, pcf)
		dfs(i, col-1, matrix, atl)
	}
	for i := 0; i < col; i++ {
		dfs(0, i, matrix, pcf)
		dfs(row-1, i, matrix, atl)
	}
	fmt.Println("###")
	fmt.Println(pcf)
	fmt.Println("###")

	fmt.Println("$$$$")
	fmt.Println(atl)
	fmt.Println("$$$$")

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if pcf[i][j] == 1 && atl[i][j] == 1 {
				res = append(res, [][]int{{i, j}}...)
			}
		}
	}
	return res
}

