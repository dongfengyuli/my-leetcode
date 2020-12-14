package main

import "fmt"

func main(){
	nums := []int{1,3,2,2,5,2,3,7}
	res := findLHS(nums)
	fmt.Println(res)
}

func findLHS(nums []int) int{
	dict := make(map[int]int,0)
	for _,v := range nums{
		dict[v]++
	}
	fmt.Println("dict",dict)
	ans := 0
	for k,v := range dict{
		if _,ok := dict[k+1];ok{
			ans = max(ans,v+dict[k+1])
		}
	}
	return ans
}



func max(a,b int) int{
	if a >= b{
		return a
	}
	return b
}

