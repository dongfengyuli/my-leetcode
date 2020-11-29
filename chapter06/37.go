package main

import "fmt"

func main(){
	n := 6
	edges := [][]int{
		{3,0},
		{3,1},
		{3,2},
		{3,4},
		{5,4},
	}

	res := findMinHeightTrees(n,edges)
	fmt.Println(res)
}


func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := make([][]int, n)
	degrees := make([]int, n)
	for _, edge := range edges {
		degrees[edge[0]] += 1
		degrees[edge[1]] += 1
		if graph[edge[0]] == nil {
			graph[edge[0]] = []int{}
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		if graph[edge[1]] == nil {
			graph[edge[1]] = []int{}
		}
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	fmt.Println("graph",graph)
	fmt.Println("degrees",degrees)

	queue := []int{}
	for i := 0; i < n; i++ {
		if degrees[i] == 1 {
			queue = append(queue, i)
		}
	}
	roots := []int{}
	for len(queue) > 0 {
		length := len(queue)
		roots = []int{}
		for i := 0; i < length; i++ {
			degrees[queue[i]] -= 1
			for _, neighbor := range graph[queue[i]] {
				degrees[neighbor] -= 1
				if degrees[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
			roots = append(roots, queue[i])
		}
		queue = queue[length:]
	}
	return roots
}