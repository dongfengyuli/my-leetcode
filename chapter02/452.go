package main

import (
	"fmt"
	"sort"
)

func main(){
	points := [][]int{{10,16},{2,8},{1,6},{7,12}}
	res := findMinArrowShots(points)

	fmt.Println(res)
}


func findMinArrowShots(points [][]int) int{
	if len(points) == 0{
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	maxRight := points[0][1]
	ans := 1

	for _,p := range points{
		if p[0] > maxRight{
			maxRight = p[1]
			ans++
		}
	}
	return ans
}