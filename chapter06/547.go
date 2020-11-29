package main

import "fmt"

func main()  {
	M := [][]int{
		{1,1,0},
		{1,1,1},
		{0,1,1},
	}
	res := findCircleNum(M)
	fmt.Println(res)
}


func findCircleNum(M [][]int) int{
	var res int
	visited := make([]int,len(M))
	for i:=0;i < len(M);i++{
		if visited[i] == 0{
			res++
			dfs(M,visited,i)
		}
	}
	return res
}


func dfs(M [][]int,visited []int,n int){
	if visited[n] == 1{
		return
	}
	visited[n] = 1
	for i:=0;i<len(M);i++{
		if M[n][i] == 1 && visited[i] == 0{
			dfs(M,visited,i)
		}
	}
	return
}