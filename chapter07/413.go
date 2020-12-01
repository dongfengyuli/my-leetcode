package main

import "fmt"

func main(){
	A := []int{1, 2, 3, 4,11,15,19}
	res := numberOfArithmeticSlices(A)

	fmt.Println(res)
}


func numberOfArithmeticSlices(A []int) int {
	return method_dp(A)
}


func method_dp(nums []int) int {
	sum := 0
	dp := make([]int, len(nums))
	for i := 2; i < len(nums); i++ {
		if nums[i] - nums[i-1] == nums[i-1] - nums[i-2] {
			dp[i] = dp[i-1] + 1
			sum += dp[i]
		}
	}
	return sum
}
