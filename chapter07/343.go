package main

import "fmt"

func main(){
	res := integerBreak(10)
	fmt.Println(res)
}



func integerBreak(n int) int{
	dp := make([]int,n+1)
	for i:=2;i<=n;i++{
		curMax := 0
		for j:=1;j<i;j++{
			curMax = max(curMax,max(j*(i-j),j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}

func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}