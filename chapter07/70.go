package main

import "fmt"

func main(){
	n := 20
	res := climbStairs(n)
	fmt.Println(res)
}


func climbStairs(n int) int{
	if 1 == n{
		return 1
	}

	f := make([]int,n+1)
	f[1],f[2] = 1,2

	for i:=3;i<len(f);i++{
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}