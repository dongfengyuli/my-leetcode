package main

import (
	"fmt"
	"sort"
)

func main (){
	intervals := [][]int{{1,2}, {2,3}, {3,4}, {1,3}}

	res := eraseOverlapIntervals(intervals)
	fmt.Println(res)
}

func eraseOverlapIntervals(intervals [][]int) int{
	if len(intervals) == 0{
		return 0
	}

	sort.Slice(intervals,func(i,j int) bool{
		if intervals[i][0] == intervals[j][0]{
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})


	fmt.Println("######")
	fmt.Println(intervals)
	fmt.Println("######")

	count := 0
	end := intervals[0][1]
	for i:=1;i < len(intervals);i++{
		if intervals[i][0] >= end{
			end = intervals[i][1]
		}else{
			if intervals[i][1] < end{
				end = intervals[i][1]
			}
			count++
		}
	}
	return count
}