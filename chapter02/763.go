package main

import "fmt"

func main(){
	S := "ababcbacadefegdehijhklij"
	res := partionLabels(S)

	fmt.Println(res)
}


func partionLabels(S string) []int{
	lastPos := [26]int{}
	partion := make([]int,0)

	for i,c := range S{
		lastPos[c-'a'] = i
	}

	start,end := 0,0
	for i,c := range S{
		if lastPos[c-'a'] > end{
			end = lastPos[c-'a']
		}
		if i == end{
			partion = append(partion,end-start+1)
			start = end + 1
		}
	}
	return partion
}