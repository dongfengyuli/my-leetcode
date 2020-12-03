package main

import "sort"

func main(){

}




func findLongestChain(pairs [][]int) int {
	//排序
	sort.Slice(pairs,func(i,j int) bool{
		if pairs[i][0] == pairs[j][0]{
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})

	n:=len(pairs)
	if (n<=1){
		return n
	}
	dp:=make([]int,n)
	for i:=0;i<n;i++{
		dp[i]=1
	}
	ans:=1
	for i:=1;i<n;i++ {
		for j:=0;j<i;j++ {
			if pairs[j][1]<pairs[i][0]{
				dp[i]=max(dp[i],dp[j]+1)
			}
		}
		ans =max(ans,dp[i])
	}
	return ans;

}
func max (x,y int) int {
	if (x<=y){
		return y
	}
	return x
}
