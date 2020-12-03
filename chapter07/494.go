package main

import "fmt"

func main(){
	nums := []int{1, 1, 1, 1, 1}
	S := 3
	res := findTargetSumWays(nums,S)
	fmt.Println(res)
}


func findTargetSumWays(nums []int, S int) int {
	sum:=0
	for e:=range nums{
		sum=sum+nums[e]
	}
	// sumAdd + sumCut = sum
	// sumAdd - sumCut = S
	// 2*sumAdd=sum+S
	// sumAdd=(sum+S)/2
	if S>sum{
		return 0
	}
	// 不是2的倍数, sumAdd不可能是小数，所以直接返回
	if (sum+S)&1!=0{
		return 0
	}
	sumAdd:=(sum+S)/2
	fmt.Println("sumAdd",sumAdd)
	// 到这里问题就转换为给定一个数组，有多少种方法装满sumAdd这个数
	dp:=make([]int,sumAdd+1)
	dp[0]=1
	for _,num:=range nums{
		for i:=sumAdd;i>=num;i--{
			dp[i]+=dp[i-num]
		}
	}
	return dp[sumAdd]
}

