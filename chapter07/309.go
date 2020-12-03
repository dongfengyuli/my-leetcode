package main

import "fmt"

func main(){
	prices := []int{1,2,3,0,2,5,9,1,4}
	res := maxProfitV2(prices)

	fmt.Println(res)
}







func maxProfitV2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2] - prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	fmt.Println("fn",f)
	return max(f[n-1][1], f[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}









//121. 买卖股票的最佳时机

//122. 买卖股票的最佳时机 II

//123. 买卖股票的最佳时机 III

//188. 买卖股票的最佳时机 IV

//309. 最佳买卖股票时机含冷冻期

//714. 买卖股票的最佳时机含手续费

//剑指 Offer 63. 股票的最大利润