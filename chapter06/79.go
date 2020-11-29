package main

import "fmt"

func main()  {
	bord := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	word := "ABCCED"

	res := exist(bord,word)
	fmt.Println(res)
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	if word == "" {
		return true
	}
	used := [][]bool{}
	for i := 0; i < m; i++ {
		used = append(used, make([]bool, n))
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == word[0] {
				used[r][c] = true
				if dfs4(board, used, r, c, word[1:]) {
					return true
				}
				used[r][c] = false
			}
		}
	}
	return false
}

func dfs4(board [][]byte, used [][]bool, row, col int, word string) bool {
	if word == "" {
		return true
	}
	ds := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var r, c int
	for _, d := range ds {
		r, c = row+d[0], col+d[1]
		if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) && !used[r][c] && board[r][c] == word[0] {
			used[r][c] = true
			if dfs4(board, used, r, c, word[1:]) {
				return true
			}
			used[r][c] = false
		}
	}
	return false
}

