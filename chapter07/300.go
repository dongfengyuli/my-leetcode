package main

import "fmt"

func main(){
	nums := []int{10,9,2,5,3,7,101,18}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}

func lengthOfLIS(nums []int) int{
	if len(nums) < 1{
		return 0
	}

	dp := make([]int,len(nums))
	result:=1
	for i:=0;i<len(nums);i++{
		dp[i] = 1
		for j:=0;j<i;j++{
			if nums[j] < nums[i]{
				dp[i] = max(dp[j] + 1,dp[i])
			}
		}
		result = max(result,dp[i])
	}
	return result
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}